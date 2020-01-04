#!bin/sh -e

#build docker image from dockerfile
docker build -t go-graphql:build -f dockerfile.build .
#create container 
docker container create --name cont-extract go-graphql:build
#copy app from container to work space
docker container cp cont-extract:/go-graphql/app ./app
#build slim image
docker build -t go-graphql:latest -f dockerfile.release .
#remove docker container
docker rm cont-extract
#remove build image
docker rmi -f go-graphql:build
rm ./app
