FROM golang:1.21 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /serve ./cmd/main.go


FROM debian:bookworm

WORKDIR /

COPY --from=build-stage /serve /serve
# COPY ./config /config
COPY ./templates /templates
COPY ./static /static

EXPOSE 8080

ENTRYPOINT ["/serve"]
