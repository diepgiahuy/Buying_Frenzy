## Build
FROM golang:1.18-alpine3.15 AS builder

ARG APP_NAME=buying-frenzy

WORKDIR /app
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o ./$APP_NAME

# Run
FROM alpine:3.15
WORKDIR /app
RUN apk add --no-cache bash
COPY --from=builder /$APP_NAME .
COPY util/config/dev.env ./util/config/dev.env
COPY migration ./migration
COPY users_with_purchase_history.json ./users_with_purchase_history.json
COPY restaurant_with_menu.json ./restaurant_with_menu.json
COPY start.sh .
EXPOSE 8080

ENTRYPOINT ["/app/start.sh"]




