FROM golang:latest

RUN mkdir /go/src/forum
WORKDIR /go/src/forum

RUN go get github.com/gorilla/mux

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go get github.com/google/wire/cmd/wire; \
    wire;

RUN go build -o main .

RUN go get github.com/pilu/fresh
RUN go get -u github.com/kardianos/govendor

EXPOSE 3010
CMD sleep 5 && fresh








