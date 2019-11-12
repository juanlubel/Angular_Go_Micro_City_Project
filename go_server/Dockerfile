FROM golang:latest

LABEL maintainer="Juanlubel <juanluis.belda@gmail.com>"

WORKDIR /go/src/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

RUN go get github.com/pilu/fresh
RUN go get -u github.com/kardianos/govendor

EXPOSE 3030

CMD ["fresh"]