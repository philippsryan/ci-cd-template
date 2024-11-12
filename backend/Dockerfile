FROM golang:1.22-alpine as builder

WORKDIR /app

COPY . ./

ENV CGO_ENABLED=0
RUN go build -v -o todoapp

FROM scratch
WORKDIR /app
COPY --from=builder /app/todoapp ./todoapp
COPY --from=builder /app/migrations ./migrations

EXPOSE 8000

CMD ["/app/todoapp", "server"]

