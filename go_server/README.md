# Go_Gingonic_Server

Endpoints

    localhost:3030/api -> "Hello":"World"
    localhost:3030/api/salute/:name -> "Hello, friend %name"}

## Dockerized

You can use a container to deploy this repository

    git clone [repository] repo_name
    cd repo_name
    docker-compose up

Start to programing GoLang without installing environment

Use ./script/govendor.sh to install dependecies on docker

## Install Go

_Before git clone_

1.- Download Go www.golang.org

2.- Download the .tar file (linux)

    tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz
    
3.- Config system routes
```
nano ~/.bashrc

~~~~
export PATH=$PATH:/user/local/go/bin
export GOROOT=/usr/local/go
export GOPATH=/home/{$user}/go
export PATH=$PATH:/home/${user}/go/bin

```
    
4.- Install GinGonic https://github.com/gin-gonic/gin#installation

    go get -u github.com/gin-gonic/gin
    go get github.com/kardianos/govendor
    
    
    govendor init
    govendor fetch route/externel/package
    
>use Govendor for production package
   
5.- Do the git clone

    mkdir ~/go/src
    cd ~/go/src
    git clone https://github.com/juanlubel/Go_Gingonic_Server.git exemple_project
    cd exemple_project
    go build
    ./exemple_project
    
*.- Import external packages

    govendor fetch github.com/route/external/package

6.- Install fresh for project auto reload

    go get github.com/pilu/fresh

Before this

    go build
    fresh
    