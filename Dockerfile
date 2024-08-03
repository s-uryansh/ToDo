FROM golang:1.22-alpine3.20

WORKDIR /CRUD-SQL

COPY go.mod go.sum ./

RUN go mod download

COPY . . 

WORKDIR /CRUD-SQL/cmd/app

RUN go build -o main .

EXPOSE 8080

CMD [ "./main" ]
