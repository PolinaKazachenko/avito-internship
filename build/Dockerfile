FROM golang:1.24.4-alpine as builder
WORKDIR /build
COPY . .

RUN go mod download
RUN go build -o /main cmd/main.go

FROM alpine:3
WORKDIR /bin
RUN apk add --no-cache make

COPY --from=builder main ./
COPY Makefile .env ./

CMD ["make", "run-in-docker"]
