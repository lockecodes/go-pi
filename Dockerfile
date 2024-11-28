FROM golang:1.23 AS base

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

FROM base AS release

COPY cmd ./cmd/
COPY configs ./configs/
COPY internal ./internal/

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /gopi /app/cmd/go-pi/main.go

# Run
CMD ["/gopi"]

FROM release AS test
COPY test ./test/

CMD ["go", "test", "/app/test/..."]