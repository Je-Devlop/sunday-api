postgres: 
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=sunday -e POSTGRES_PASSWORD=secret -d postgres:15-alpine

image:
	docker build -t sunday:test -f Dockerfile .

container:
	docker run -p:8081:8081 -e DB_CONN='postgres://sunday:secret@db:5432/sunday' --link postgres15:db \
	--name sunday sunday:test 

.PHONY: postgres image container
 