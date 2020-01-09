FROM golang:latest as builder

RUN mkdir /build
ADD ./ /build 
#ver la direccion

WORKDIR /build
RUN env GOOS=linux GOARCH=386 go build -o main .

FROM alpine:latest

RUN apk update \
        && apk upgrade \
        && apk add --no-cache \
        ca-certificates \
        && update-ca-certificates 2>/dev/null || true

RUN mkdir -p /app && adduser -S -D -H -h /app appuser && chown -R appuser /app
RUN mkdir /app/config
COPY --from=builder /build/config.toml /app/
COPY --from=builder /build/main /app/

USER appuser
EXPOSE 9091
WORKDIR /app
CMD ["./main"]