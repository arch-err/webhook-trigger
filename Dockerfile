#!CMD: make build && make run
FROM golang:1.23 AS builder

COPY ./app /build
WORKDIR /build

RUN go mod download

ARG BUILDARGS=""
ENV CGO_ENABLED=0
ENV BUILDENV="GOOS=linux"

RUN go build ${BUILDARGS} -ldflags '-extldflags "-static"' -o webzippy

# ---

FROM alpine:3.22.1

LABEL org.opencontainers.image.source https://github.com/arch-err/webhook-trigger

COPY --from=builder /build/webhook-trigger /
COPY --from=builder /build/static /static
COPY --from=builder /build/views /views
COPY --from=builder /build/archives /archives

EXPOSE 80
ENTRYPOINT ["/webhook-trigger"]
