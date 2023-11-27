FROM golang:1.21-alpine
RUN apk update && apk upgrade && \
    apk add  build-base --no-cache bash git openssh
WORKDIR /app
ENV CGO_ENABLED=1
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .
CMD ["./main"]