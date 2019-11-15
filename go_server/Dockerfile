FROM golang:latest

LABEL maintainer="Juanlubel <juanluis.belda@gmail.com>"

#RUN mkdir /go/src/go_server
WORKDIR /go/src/go_server

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

RUN go get github.com/pilu/fresh; \
    go get -u github.com/kardianos/govendor

EXPOSE 3030

CMD ["fresh"]
#CMD ./main