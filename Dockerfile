FROM golang:1.21.3

WORKDIR /app
COPY . ./
COPY config/config.yml ./cmd/api/config/
RUN go mod download

WORKDIR /app/cmd/api
RUN go build -o /app-api

EXPOSE 8080
CMD ["/app-api"]
