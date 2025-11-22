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
COPY --from=builder /app/tracker .
COPY --from=builder /app/scaffold .
COPY --from=builder /app/scripts/run_update.sh .
COPY --from=builder /app/entrypoint.sh .

# Make sure scripts are executable
RUN chmod +x entrypoint.sh run_update.sh

# Create a place for git config
RUN mkdir -p /root/.ssh

ENTRYPOINT ["./entrypoint.sh"]
