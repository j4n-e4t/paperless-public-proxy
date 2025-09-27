FROM docker.io/golang:1.25.1-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main .

# Start a new stage using alpine for ca-certificates
FROM alpine:latest AS certs

FROM scratch

# Copy the ca-certificates from the certs stage
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /app/main /app/main

ENV PAPERLESS_BASE_URL=http://paperless:8000/
ENV TS_HOSTNAME=paperless-public-proxy
ENV TS_AUTHKEY=your-key
ENV TS_DIR=/app/tsnet-data

CMD [ "/app/main" ]