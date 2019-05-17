# Go parameters
    GOCMD=go
    GOBUILD=$(GOCMD) build
    GOCLEAN=$(GOCMD) clean
    GOTEST=$(GOCMD) test
    GOGET=$(GOCMD) get
    BINARY_NAME=dronedelivery
	BINARY_PATH=./cmd/dronedelivery/
    INPUT_FILE_PATH=./assets/files/orders.txt
    OUTPUT_FILE_PATH=./assets/files/schedule.txt
    DOCKER_IMAGE=anup8000/dronedelivery:0.1
    DOCKERFILE=./build/package/dockerfile
    DOCKERBUILD=docker build
    DOCKERRUN=docker run

    
    all: test build
    install:
			$(GOBUILD) -o $(BINARY_PATH)$(BINARY_NAME) $(BINARY_PATH)
    test:
			$(GOTEST) -v ./internal/pkg/fileops/
			$(GOTEST) -v ./internal/app/schedule/
    clean:
			$(GOCLEAN)
			rm -f $(BINARY_PATH)$(BINARY_NAME)
    run:
			$(GOBUILD) -o $(BINARY_PATH)$(BINARY_NAME) $(BINARY_PATH)
			$(BINARY_PATH)$(BINARY_NAME) -i $(INPUT_FILE_PATH) -o $(OUTPUT_FILE_PATH)
