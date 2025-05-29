FROM docker.io/golang:1.24.2-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main .

FROM scratch

COPY --from=builder /app/main /app/main

ENV PAPERLESS_BASE_URL=http://paperless:8000/share/
ENV PROXY_URL=http://localhost:8080

EXPOSE 8080

CMD [ "/app/main" ]