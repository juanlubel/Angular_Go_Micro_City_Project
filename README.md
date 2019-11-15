#Angular Go app

##Generic

Docker environment in a global .env
    
* app-variables -> client
* db.env -> mysql

##Angular

_no dockerized_

##Go

Install dependencies in go_server docker

>_Using a shell script (exemples bellow):_

    ./script/govendor.sh fetch route/to/package
    ./script/govendor.sh remove route/to/package
    ./script/govendor.sh list

Endpoints


    http://localhost:3030/api -> "Hello world"
    http://localhost:3030/api/salute/:name -> Hello, my friend ${name}

##Mysql / Adminer