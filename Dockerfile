FROM golang:1.12.0-alpine3.9 as builder
RUN mkdir -p /go/src/github.com/fjukstad/headwind
ADD . /go/src/github.com/fjukstad/headwind
WORKDIR /go/src/github.com/fjukstad/headwind
RUN go build -o /app/headwind .
RUN cp /go/src/github.com/fjukstad/headwind/index.html /app/index.html

FROM alpine:3.11.3
RUN apk add ca-certificates
COPY --from=builder /app /app
WORKDIR /app
CMD ["/app/headwind"]
