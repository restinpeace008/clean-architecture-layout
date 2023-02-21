MYDIR=bin
run:
	@mkdir -p $(MYDIR) && go build -o $(MYDIR)/app ./cmd/app/ && ./bin/app