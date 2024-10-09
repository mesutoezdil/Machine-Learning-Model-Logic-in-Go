FROM golang:1.22-alpine
WORKDIR /app
COPY go.mod .
COPY . .
RUN go mod download
RUN go build -o model-app .
CMD ["/app/model-app"]
