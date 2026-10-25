FROM golang:tip-alpine3.24

# Set destination for COPY
WORKDIR /app

# Enable dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY ./internal ./internal
COPY ./cmd ./cmd

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /mealgen ./cmd

EXPOSE 8080

# Run
CMD ["/mealgen"]
