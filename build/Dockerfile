FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates
RUN adduser -D -g '' appuser

WORKDIR $GOPATH/src/turtle-wear-api

COPY . .

RUN go mod download && go mod verify
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go test -cover ./...
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /go/bin/turtle-wear-api ./cmd/turtle-functions/main.go

FROM scratch
COPY --from=builder /etc/ssl/certs /etc/ssl/certs
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /go/bin/turtle-wear-api /go/bin/turtle-wear-api
USER appuser
ENTRYPOINT ["/go/bin/turtle-wear-api"]