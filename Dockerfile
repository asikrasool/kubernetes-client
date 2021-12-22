FROM golang:latest as builder
# Define build env
ARG USE_KUBECONFIG
ENV USE_KUBECONFIG=${USE_KUBECONFIG}
# Add a work directory
WORKDIR /app
# Cache and install dependencies
COPY go.mod go.sum ./
RUN go mod download
# Copy app files
COPY . .
# Build app
RUN go build -o app
ENTRYPOINT ["./app"]


# FROM alpine as production
# # Add certificates
# RUN apk add --no-cache ca-certificates
# # Copy built binary from builder
# WORKDIR /app

# COPY --from=builder app .

# # Exec built binary
# ENTRYPOINT ["./app"]