FROM golang:1.23.1-alpine
LABEL authors="yoshutch"

WORKDIR /app

COPY . ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o run-app

EXPOSE 8080

# Run
CMD ["./run-app"]