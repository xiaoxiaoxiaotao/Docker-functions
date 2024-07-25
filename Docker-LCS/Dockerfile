FROM golang:1.18.1

WORKDIR /Docker_gin_functions

COPY . .

RUN go mod download

RUN go mod tidy

EXPOSE 5000

CMD [ "go", "run" , "app.go"]