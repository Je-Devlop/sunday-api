postgres: 
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=sunday -e POSTGRES_PASSWORD=secret -d postgres:15-alpine

.PHONY: postgres
 