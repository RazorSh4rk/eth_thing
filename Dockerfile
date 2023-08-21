FROM golang:1.21 AS builder

WORKDIR /app
COPY . .
RUN go build -o eth_tracker

# Use a smaller base image for the runtime container
# FROM debian:buster-slim

# WORKDIR /app
# COPY --from=builder /app/eth_tracker .
EXPOSE 8080
# RUN apt install libc6

CMD ["./eth_tracker"]
