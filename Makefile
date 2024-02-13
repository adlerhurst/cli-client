generate-option:
	$(RM) -r protoc-gen-cli-client/.artifacts
	cd protoc-gen-cli-client && buf generate 
	mkdir -p protoc-gen-cli-client/option
	mv protoc-gen-cli-client/.artifacts/adlerhurst/cli/v1alpha/* protoc-gen-cli-client/option
	$(RM) -r protoc-gen-cli-client/.artifacts

.PHONY: compile
compile: generate-option
	cd protoc-gen-cli-client && go mod tidy
	cd protoc-gen-cli-client && go install

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