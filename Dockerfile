ARG stage

# stage - golang app
FROM golang:1.17-alpine3.14 as builder

WORKDIR /app
ADD . /app/

#RUN apk update && apk add git
RUN apk add build-base
RUN go mod download
RUN go get github.com/githubnemo/CompileDaemon
RUN GOOS=linux go build -o /app/main

### PROD
FROM alpine:3.14 as go-prod
WORKDIR /app
COPY --from=builder /app/main /app/gobot

## Healthcheck for PROD
#RUN apk --no-cache add curl
#HEALTHCHECK --interval=5s --timeout=3s --start-period=10s --retries=3 \
#    CMD curl -fs http://localhost:$PORT/ || exit 1

ENTRYPOINT ./gobot


### DEV
FROM builder as go-dev
ENTRYPOINT CompileDaemon -directory=/app -build="go build main.go" -command=./main

FROM go-${stage} AS final


