#First Stage
FROM golang:latest AS builder

RUN mkdir /app

ADD . /app

RUN CGO_ENABLED=0 GOOS=linux go build -o csgo-exporter 


#Second Stage
FROM alpine:latest AS production

COPY --from=builder /app .

CMD["./csgo-exporter"]
