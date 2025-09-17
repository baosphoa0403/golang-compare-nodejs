# Sử dụng Golang image chính thức từ Docker Hub làm base image
FROM golang:1.24

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /main 


EXPOSE 8888

CMD ["/main"]
