FROM golang
WORKDIR /app
COPY . .
RUN go mod tidy
CMD go run main.go postgresql://root@roach_alb:26257/defaultdb?sslmode=disable