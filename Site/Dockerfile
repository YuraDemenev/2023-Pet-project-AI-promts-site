FROM golang:1.21.0

RUN go version

# WORKDIR /app
# COPY go.mod go.sum ./
# RUN go mod download

# COPY *.go ./
# COPY *.yml ./
# COPY *.env ./
# COPY *.jpg ./

# RUN go build -o site ./cmd/main.go
# EXPOSE 8080
# CMD ["./site"]

COPY ./config/ /config/
# COPY /cmd/.env /cmd

ENV GOPATH=/cmd
COPY ./ ./

RUN go mod download
RUN go build -o site ./cmd/main.go

CMD ["./site"]