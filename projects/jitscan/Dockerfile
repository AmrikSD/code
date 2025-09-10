FROM --platform=$BUILDPLATFORM golang:1.24-alpine AS builder
ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o jitscan .

FROM gcr.io/distroless/static:nonroot
COPY --from=builder /app/jitscan /jitscan
EXPOSE 8443
USER 65532:65532
ENTRYPOINT ["/jitscan"]
