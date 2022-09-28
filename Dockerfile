# build stage
FROM golang:1.19-alpine as builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64


WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o fswatch-command .

# runtime stage
FROM alpine
COPY --from=builder /build/fswatch-command /app/
RUN chmod +x /app/fswatch-command
CMD ["/app/fswatch-command"]
