FROM golang:rc-alpine3.13 AS builder
WORKDIR /app

COPY . .
COPY go.mod .
COPY go.sum .

RUN go clean --modcache
RUN GOOS=linux go build -o notifier notifier/notifier.go

FROM alpine:3.13
WORKDIR /app

RUN apk --no-cache add tzdata

ENV TZ=Asia/Jakarta

COPY --from=builder /app/notifier .
COPY ["notifier/appsettings.Dev.json","."]
COPY ["notifier/appsettings.Prod.json","."]
COPY ["notifier/appsettings.Stag.json","."]

CMD [ "/app/notifier" ]