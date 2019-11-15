FROM golang:latest

LABEL maintainer="Juanlubel <juanluis.belda@gmail.com>"

<<<<<<< HEAD:go_server/Dockerfile
#RUN mkdir /go/src/go_server
WORKDIR /go/src/go_server
=======
RUN mkdir /go/src/idcards
WORKDIR /go/src/idcards
>>>>>>> feature-traefik:idcards/Dockerfile

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

RUN go get github.com/pilu/fresh; \
    go get -u github.com/kardianos/govendor

EXPOSE 3030

<<<<<<< HEAD:go_server/Dockerfile
CMD ["fresh"]
#CMD ./main
=======
ENTRYPOINT ["fresh"]
>>>>>>> feature-traefik:idcards/Dockerfile
