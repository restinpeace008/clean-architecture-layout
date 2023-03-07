MYDIR=bin
serve: build run

build: 
	@mkdir -p $(MYDIR) && go build -o $(MYDIR)/app ./cmd/app/

.PHONY: build

run:
	@./bin/app