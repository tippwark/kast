FROM golang:1.27-alpine AS builder

ARG KAST_VERSION

WORKDIR /src
ADD . .
RUN go build -ldflags "-X \"main.kastVersion=${KAST_VERSION}\" -X \"main.compileDate=$(date)\"" -o ./bin/kast ./cmd/kast

# ---

FROM debian:latest AS kubectl_downloader

RUN apt-get update && apt-get install -y ca-certificates curl && rm -rf /var/lib/apt/lists/*
RUN curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
RUN chmod +x ./kubectl

# ---

FROM debian:latest

RUN apt-get update && apt-get install -y curl xz-utils && rm -rf /var/lib/apt/lists/*
COPY --from=builder /src/bin/kast /usr/bin/kast
COPY --from=kubectl_downloader /kubectl /usr/bin/kubectl
