# build stage
FROM golang:1.14-alpine AS builder
WORKDIR /app
COPY . .

RUN apk add --no-cache git
RUN go build -v -o app .

# final stage
FROM alpine:latest
WORKDIR /root
COPY ./.env.development /
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/app .
ENTRYPOINT ENV=DEV ./app
