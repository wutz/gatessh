FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o gatessh-controller ./cmd/controller/

FROM alpine:3.21
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/gatessh-controller /usr/local/bin/
ENTRYPOINT ["gatessh-controller"]
