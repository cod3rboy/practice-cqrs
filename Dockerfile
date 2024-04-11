FROM golang:1.22.2-alpine AS build

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o api cmd/server/main.go

FROM alpine:3.18

WORKDIR /

COPY --from=build /app/api /usr/bin/

EXPOSE 3001

ENTRYPOINT api
