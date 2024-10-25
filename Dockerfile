ARG GO_VERSION=1
FROM golang:${GO_VERSION}-bookworm as builder

RUN apt-get update && apt-get install -y \
    nodejs npm

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY package.json package-lock.json ./
RUN npm install
COPY . .
RUN go build -v -o /run-app ./cmd/server

FROM debian:bookworm

RUN apt-get update && apt-get install -y \
    ca-certificates && \
    apt-get clean

WORKDIR /usr/src/app
COPY --from=builder /run-app /usr/local/bin/
COPY --from=builder /usr/src/app/config ./config/
COPY --from=builder /usr/src/app/internal/translations ./internal/translations/ 
COPY --from=builder /usr/src/app/internal/templates ./internal/templates/ 

CMD ["run-app"]
