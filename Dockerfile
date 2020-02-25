#First Stage
FROM golang:latest AS builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o csgoexporter 

#Second Stage
FROM alpine:latest AS production

COPY --from=builder /app/csgoexporter .

CMD ["./csgoexporter"]
