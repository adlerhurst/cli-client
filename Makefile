generate-option:
	$(RM) -r protoc-gen-go-cli/.artifacts
	cd protoc-gen-go-cli && buf generate 
	mkdir -p protoc-gen-go-cli/option
	mv protoc-gen-go-cli/.artifacts/adlerhurst/cli/v1alpha/* protoc-gen-go-cli/option
	$(RM) -r protoc-gen-go-cli/.artifacts

.PHONY: compile
compile: generate-option
	cd protoc-gen-go-cli && go mod tidy
	cd protoc-gen-go-cli && go install

generate-example:
	$(RM) -r example/proto/*pb.go
	cd example && buf generate
	mkdir -p example/cli
	mv example/.artifacts/adlerhurst/example/v1/* example/cli
	$(RM) -r example/.artifacts

.PHONY: compile-example
compile-example: compile generate-example
	cd example && go mod tidy
	cd example && go install