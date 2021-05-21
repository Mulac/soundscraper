FROM golang:1.16
EXPOSE 8080

WORKDIR /app

# Cache dependancies
COPY go.mod go.sum ./
RUN go mod download

# Copy source and build
COPY . .
RUN go build -o main .

ENTRYPOINT ["./main"]