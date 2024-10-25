ARG GO_VERSION=1
FROM golang:${GO_VERSION}-bookworm as builder

RUN apt-get update && apt-get install -y \
    nodejs npm \
    python3 make g++ gcc

WORKDIR /usr/src/app
COPY go.mod go.sum ./
COPY package.json package-lock.json ./
RUN go mod download && go mod verify
RUN npm cache clean --force && npm install --legacy-peer-deps --verbose
RUN go build -v -o /run-app ./cmd/server

FROM debian:bookworm

RUN apt-get update && apt-get install -y \
    ca-certificates

COPY --from=builder /run-app /usr/local/bin/

WORKDIR /usr/src/app
COPY . .

CMD ["run-app"]
