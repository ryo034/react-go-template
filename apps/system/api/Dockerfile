FROM golang:1.21.5-alpine as builder

ARG GROUPNAME=user
ARG USERNAME=user
ARG GID=1001
ARG UID=1001
RUN addgroup -S -g $GID $GROUPNAME && adduser -u $UID -S $USERNAME -G $GROUPNAME

# For Import Private Package
ARG GITHUB_TOKEN

RUN apk update && \
    apk add --no-cache git ca-certificates tzdata && \
    git config --global url."https://${GITHUB_TOKEN}:x-oauth-basic@github.com/ryo034/".insteadOf "https://github.com/ryo034/"

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies using go modules.
# Allows container builds to reuse downloaded dependencies.
COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY ../.. ./

# Build the binary.
# -mod=readonly ensures immutable go.mod and go.sum in container builds.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=readonly -buildvcs=false -v -o server

FROM scratch

# Expose 8080port is required when running in Cloud Run.
EXPOSE 8080

# Copy ca-certificates.crt
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Need Timezone for Golang time package
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/server /server

USER $USERNAME

# Run the web service on container startup.
CMD ["/server"]
