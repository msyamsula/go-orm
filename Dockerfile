FROM golang:latest

WORKDIR /go-orm

COPY ./connection connection
COPY ./jsonRequest jsonRequest
COPY ./jsonResponse jsonResponse
COPY ./main main
COPY ./model model
COPY ./service service
COPY ./go.mod go.mod

RUN go mod tidy

EXPOSE 5000

CMD [ "go", "run", "main/main.go" ]
