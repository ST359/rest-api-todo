FROM golang:1.23

WORKDIR ${GOPATH}/todo-app/
COPY . ${GOPATH}/todo-app/
RUN go build -o /build ./cmd/todo-app/ \
    && go clean -cache -modcache

EXPOSE 8080

CMD ["/build"]