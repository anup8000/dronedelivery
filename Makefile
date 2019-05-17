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

    
    all: test build
    install:
			$(GOBUILD) $(BINARY_PATH)
    test:
			$(GOTEST) -v ./internal/pkg/fileops/
			$(GOTEST) -v ./internal/app/schedule/
    clean:
			$(GOCLEAN)
			rm -f $(BINARY_PATH)$(BINARY_NAME)
    run:
			$(GOBUILD) $(BINARY_PATH)
			$(BINARY_PATH)$(BINARY_NAME) -i $(INPUT_FILE_PATH) -o $(OUTPUT_FILE_PATH)