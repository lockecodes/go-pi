FROM golang:1.23

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY api ./api/
COPY assets ./assets/
COPY cmd ./cmd/
COPY configs ./configs/
COPY init ./init/
COPY internal ./internal/
COPY pkg ./pkg/
COPY vendor ./vendor/
COPY web ./web/
COPY website ./website/

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /my-app

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 8080

# Run
CMD ["/my-app"]
