GO_PACKAGES=$(shell go list ./... | grep -v vendor)
GOLANGCI_VERSION=v1.51.0

lint:
	golangci-lint run

# Install golangci-lint	
install-lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ${GO_PATH}/bin ${GOLANGCI_VERSION}

vet:
	go vet ${GO_PACKAGES}