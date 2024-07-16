FROM golang:1.18-buster AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o ./audit ./cmd/audit/main.go

FROM debian:buster-slim

WORKDIR /opt/bin/
RUN apt-get update \
 && apt-get install -y --no-install-recommends ca-certificates

RUN update-ca-certificates
COPY --from=builder /src/audit /opt/bin/audit
ENV ENV_CONFIG_ONLY=true
CMD ["/opt/bin/audit"]
