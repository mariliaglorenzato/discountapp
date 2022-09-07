FROM golang:1.18 AS build

LABEL maintainer="<marilia>"

# create a working directory inside the image
WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

ENV PORT 8080
EXPOSE 8080

RUN go build 

CMD ["./discountapp"]
