FROM golang:latest as builder

RUN mkdir /build
ADD /colegios-session /build

RUN mkdir -p /root/.ssh

# Absolute source path RELATIVE TO CONTEXT so we should have physically the ssh directory as a copy of our ~/.ssh
RUN chmod 700 /root/.ssh/id_rsa


WORKDIR /build

RUN env GOOS=linux GOARCH=386 go build -o main .

FROM alpine:latest

RUN mkdir -p /app && adduser -S -D -H -h /app appuser && chown -R appuser /app
COPY --from=builder /build/main /build/config.toml /app/
USER appuser
EXPOSE 9091
WORKDIR /app
CMD ["./main"]