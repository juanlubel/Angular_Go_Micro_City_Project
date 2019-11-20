FROM golang:latest

LABEL maintainer="Juanlubel <juanluis.belda@gmail.com>"


RUN mkdir /go/src/idcards
WORKDIR /go/src/idcards


COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

RUN go get github.com/pilu/fresh; \
    go get -u github.com/kardianos/govendor

EXPOSE 3030


ENTRYPOINT ["fresh"]

