# build golang
FROM golang:1.27-alpine AS builder

WORKDIR /src

# download dependencies 
COPY go.mod go.sum ./
RUN go mod download

# copy source code
COPY cmd ./cmd
COPY internal ./internal

# linux build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build \
    -trimpath \
    -ldflags="-s -w" \
    -o /portfolio \
    ./cmd/portfolio


# runtime 
FROM alpine:3.22

# create unprivileged user
RUN addgroup -S portfolio && \
    adduser -S -G portfolio portfolio

WORKDIR /app

# copy compiled binary
COPY --from=builder /portfolio /app/portfolio

# mount directory for SSH host key
RUN mkdir /data && \
    chown portfolio:portfolio /data

USER portfolio

# port
EXPOSE 42069

ENTRYPOINT ["/app/portfolio"]
