#!/bin/bash

# docker pull mysql:5.7

docker run --name mysql -p 3306:3306 -v $PWD/config:/etc/mysql/conf.d -v $PWD/data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=test -d mysql:5.7
