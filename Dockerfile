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

FROM golang:1.27-alpine AS licenses

WORKDIR /src
ADD . .
RUN go run github.com/google/go-licenses@latest save ./... --save_path=THIRD_PARTY_LICENSES

# ---

FROM debian:latest AS final

RUN apt-get update && apt-get install -y curl xz-utils && rm -rf /var/lib/apt/lists/*

COPY --from=licenses /src/THIRD_PARTY_LICENSES /usr/share/doc/kast/THIRD_PARTY_LICENSES
COPY --from=builder /src/LICENSE /usr/share/doc/kast/LICENSE

COPY --from=builder /src/bin/kast /usr/bin/kast
COPY --from=kubectl_downloader /kubectl /usr/bin/kubectl
