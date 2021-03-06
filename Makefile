# Heavily based off https://sohlich.github.io/post/go_makefile/
# pared down from the above and added deps download via go mod

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GETDEPS=$(GOCMD) mod download
BINARY_NAME=apkreport

all: deps build
install: $(BINARY_NAME)
		mv $(BINARY_NAME) /usr/local/bin/$(BINARY_NAME)
build:
		$(GOBUILD) -o $(BINARY_NAME) -v
		@echo "build done run with: ./$(BINARY_NAME)"
		@echo "or install with 'sudo make install' to install to /usr/local/bin/$(BINARY_NAME)"
test:
		$(GOTEST) -v ./...
clean:
		$(GOCLEAN)
		rm -f $(BINARY_NAME)
run:
		$(GOBUILD) -o $(BINARY_NAME) -v
		./$(BINARY_NAME)
deps:
		$(GOGETGETDEPS)
