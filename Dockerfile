FROM golang:1.15.2 AS builder

WORKDIR /go/src/github.com/ibreakthecloud/contact-book
# Copy all files
COPY . .
# Enable Go Modules
ENV GO111MODULE=on
# Fetch dependencies before go build
RUN go mod download
# Build the binary
RUN CGO_ENABLED=1 go build -o cbook

FROM ubuntu AS final

COPY --from=builder /go/src/github.com/ibreakthecloud/contact-book/cbook /
ENV PORT 8080
EXPOSE 8080
CMD ["/cbook"]