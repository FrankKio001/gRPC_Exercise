# PowerShell
BIN_DIR = bin
PROTO_DIR = HelloWorld/proto
SERVER_DIR = HelloWorld/server
CLIENT_DIR = HelloWorld/client

HelloWorld: ## Generate Pbs and build for HelloWorld
	@echo "Building protobuf files..."
	@PowerShell -NoProfile -Command "protoc -I$(PROTO_DIR) $(PROTO_DIR)/*.proto --go_out=$(PROTO_DIR) --go-grpc_out=$(PROTO_DIR) --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative"
	@echo "Building server and client..."
	@PowerShell -NoProfile -Command "go build -o $(BIN_DIR)/HelloWorld/server.exe ./$(SERVER_DIR)"
	@PowerShell -NoProfile -Command "go build -o $(BIN_DIR)/HelloWorld/client.exe ./$(CLIENT_DIR)"

clean_hw: ## Clean generated files for HelloWorld
	@echo "Cleaning generated files..."
	@PowerShell -NoProfile -Command "Remove-Item -Path '$(PROTO_DIR)/*.pb.go' -Force"

help: ## Show this help
	@echo "Available commands:"
	@echo "  make HelloWorld - Build the HelloWorld application, including protobuf files."
	@echo "  make clean_hw - Clean the generated files."
	@echo "  make help - Display this help message."

about: ## Display info related to the build
	@echo "OS: ${OS}"
	@echo "Shell: ${SHELL} ${SHELL_VERSION}"
	@PowerShell -NoProfile -Command "protoc --version"
	@PowerShell -NoProfile -Command "go version"
	@PowerShell -NoProfile -Command "echo 'Go package: ${PACKAGE}'"
	@PowerShell -NoProfile -Command "openssl version"

.PHONY: HelloWorld clean_hw help
