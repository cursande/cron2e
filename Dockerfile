FROM golang:1.12-alpine as builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify
COPY . .

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /app/cron2e

###########################
FROM scratch

COPY --from=builder /app/cron2e ./cron2e

ENTRYPOINT ["./cron2e"]
