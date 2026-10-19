FROM golang:1.27-alpine
WORKDIR /src
COPY go.mod go.sum main.go ./
RUN go build -o app .
EXPOSE 8080
CMD ["./app"]