FROM golang:1.7.3 AS builder
WORKDIR /app
COPY hueControl.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o hueControl .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /app/
COPY --from=builder /app .
ENTRYPOINT ["./hueControl"]