FROM golang:1.23 AS builder

WORKDIR /app
COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /todo-app ./cmd/todo-app/ \
    && go clean -cache -modcache

FROM alpine:latest
WORKDIR /
COPY --from=builder /todo-app ./todo-app
RUN ls -l

EXPOSE 8080

CMD ["/todo-app"]