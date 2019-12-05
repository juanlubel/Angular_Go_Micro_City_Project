# Start from golang:1.13-alpine base image
FROM golang:1.13-alpine

# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
#RUN apk update && apk upgrade && \
#    apk add --no-cache bash git openssh

RUN mkdir /go/src/redis
WORKDIR /go/src/redis


COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go get github.com/google/wire/cmd/wire
    
RUN go build -o main .

RUN go get github.com/pilu/fresh
RUN go get -u github.com/kardianos/govendor

EXPOSE 3015
CMD sleep 5 && fresh