FROM golang:alpine AS builder

WORKDIR /app
COPY . .
RUN go build -o tracker ./cmd/tracker
RUN go build -o scaffold ./cmd/scaffold

FROM alpine:latest

# Install git, openssh, and tzdata
RUN apk add --no-cache git openssh-client tzdata

# Set Timezone to Oslo
ENV TZ=Europe/Oslo

WORKDIR /app

# Copy binaries to /usr/local/bin
COPY --from=builder /app/tracker /usr/local/bin/tracker
COPY --from=builder /app/scaffold /usr/local/bin/scaffold

# Copy scripts to /usr/local/bin
COPY --from=builder /app/scripts/run_update.sh /usr/local/bin/run_update.sh
COPY --from=builder /app/docker/entrypoint.sh /usr/local/bin/entrypoint.sh

# Make sure scripts are executable
RUN chmod +x /usr/local/bin/entrypoint.sh /usr/local/bin/run_update.sh

# Create a place for git config
RUN mkdir -p /root/.ssh

ENTRYPOINT ["entrypoint.sh"]
