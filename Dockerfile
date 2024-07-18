FROM golang:1.18 as builder

WORKDIR /app

COPY . .
RUN go mod download
RUN go build -o main .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

RUN chmod +x ./main

ENV DATABASE_URL=postgres://user:password@db/project-managment?sslmode=disable

CMD ["./main"]
