#!bin/sh -e

#build docker image from dockerfile
docker build -t go-graphql:build -f dockerfile .
#create container 
docker container create --name cont-extract go-graphql:build
#copy app from container to work space
docker container cp cont-extract:/go-graphql/app ./app
#build slim image
docker build -t go-graphql:latest -f dockerfile.release .
