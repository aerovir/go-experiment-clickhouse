DOCKER=docker-compose --project-directory deployments -f deployments/docker-compose.yml

help: Makefile
	@echo "Choose a command run:"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'

## up: Start application in docker containers with hot reload
up:
	@$(DOCKER) up -d --build

## down: Stop application and remove docker containers
down:
	@$(DOCKER) down --remove-orphans

## lint: Check source code by linters
lint:
	@echo "Checking go vet..." && go vet ./... && echo "Done!\n"
	@echo "Checking golangci-lint..." && golangci-lint run ./... && echo "Done!"
