MYDIR=bin
serve: build run

build: 
	@mkdir -p $(MYDIR) && go build -o $(MYDIR)/app ./cmd/app/

.PHONY: build

run:
	@./bin/app

up: 
	docker compose -f $(PWD)/build/package/docker-compose.yml up -d --build

down: 
	docker-compose -f $(PWD)/build/package/docker-compose.yml down