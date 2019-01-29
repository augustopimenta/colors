FROM golang:1.11 as builder
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o colors

FROM scratch
WORKDIR /app
COPY --from=builder /src/colors .
COPY --from=builder /src/index.html .
ENTRYPOINT ["/app/colors"]