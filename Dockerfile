FROM alpine:latest
RUN mkdir -p /app
WORKDIR /app

ADD main /app

CMD ["./main"]
