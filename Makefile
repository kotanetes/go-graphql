db:
	docker pull mongo:latest
	docker run --name local-mongo -p 27017:27017 -d mongo:latest

db-up:
	mongo 127.0.0.1/school script.js

db-down:
	docker stop local-mongo
	docker rm -f local-mongo