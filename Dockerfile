FROM golang:1.15-alpine3.12 as builder
RUN apk add --no-cache alpine-sdk
WORKDIR /build
COPY [ ".", "." ]
RUN make build

FROM alpine:3.12 as release
RUN apk --no-cache add \
  ca-certificates=20191127-r4 \
  bash=5.0.17-r0 \
  bash-completion=2.10-r0 && \
  echo "source <(loli completion bash)" >> ~/.bashrc
COPY --from=builder [ "/build/bin/loli", "/usr/local/bin/loli" ]
RUN chmod +x /usr/local/bin/loli
CMD ["loli"]
