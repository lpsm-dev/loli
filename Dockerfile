FROM golang:alpine as builder
WORKDIR /build
COPY [ ".", "." ]
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-extldflags "-static"' -o ./bin/loli main.go

FROM alpine:3.12 as release
COPY --from=builder [ "/build/bin/loli", "/usr/local/bin/loli" ]
RUN chmod +x /usr/local/bin/loli
RUN apk --no-cache add \
  ca-certificates=20191127-r4 \
  bash=5.0.17-r0
CMD ["loli"]
