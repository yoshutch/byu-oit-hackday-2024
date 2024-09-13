FROM golang:1.19.13-alpine
LABEL authors="yoshutch"

WORKDIR /app

COPY go.mod ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

EXPOSE 8080

# Run
CMD ["/docker-gs-ping"]