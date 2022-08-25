FROM golang:latest
WORKDIR /app

COPY . .

RUN go mod download
RUN apt-get update && apt-get install build-essential -y
RUN go get -u gorm.io/gorm
RUN go get -u gorm.io/driver/sqlite
RUN go build -o run-app

EXPOSE 1323

CMD [ "./run-app" ]
