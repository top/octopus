FROM golang:alpine AS builder
RUN mkdir -p /build
WORKDIR /build
ADD . /build
RUN CGO_ENABLED=1 go build

FROM alpine:edge
RUN mkdir -p /app
WORKDIR /app
COPY --from=builder /build/octopus .
COPY --from=builder /build/config.yaml .
EXPOSE 80
CMD [ "./octopus" ]
