define ASCII_ART
             _       _         
   ___ _   _| | __ _| |__  ___ 
  / _ \ | | | |/ _` | '_ \/ __|
 |  __/ |_| | | (_| | |_) \__ \

  \___|\__,_|_|\__,_|_.__/|___/   

   simple makefile | run `make install` to execute application locally

endef
export ASCII_ART

default: help

.PHONY: help
help: # Show all commands available to execution 
	@echo "$$ASCII_ART"
	@grep -E '^[a-zA-Z0-9 -]+:.*#'  Makefile | while read -r l; do printf "\033[1;34m$$(echo $$l | cut -f 1 -d':')\033[00m:$$(echo $$l | cut -f 2- -d'#')\n"; done

.PHONY: install
install: # Install eulabs application
	docker-compose up -d --build
	@echo "\033[1;34mContainers installed with success, follow the steps below to finish!\033[00m"
	go run ./cmd/eulabs/main.go

.PHONY: run
run: # Command to locally run application 
	go run ./cmd/eulabs/main.go

.PHONY: test
test: # Run tests eulabs application
	go test ./...

.PHONY: clean
clean: # Clean eulabs application (drop all containers)
	docker-compose down

.PHONY: env
env: # Replace .env.example to .env
	cp .env.example .env