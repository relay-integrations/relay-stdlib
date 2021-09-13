FROM golang:1.16-alpine AS builder
ENV CGO_ENABLED 0
WORKDIR /build
COPY . .
RUN go build -a -o /bin/go-template ./cmd/go-template

FROM alpine:latest
RUN apk --no-cache add ca-certificates && update-ca-certificates
COPY --from=builder /bin/go-template /bin/go-template
ENTRYPOINT []
CMD ["/bin/go-template"]
