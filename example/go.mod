module github.com/adlerhurst/cli-client/example

go 1.22.0

require (
	github.com/adlerhurst/cli-client v0.0.0-00010101000000-000000000000
	github.com/spf13/pflag v1.0.5
	google.golang.org/grpc v1.61.0
	google.golang.org/protobuf v1.32.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	golang.org/x/net v0.18.0 // indirect
	golang.org/x/sys v0.14.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231106174013-bbf56f31fb17 // indirect
)

require github.com/spf13/cobra v1.8.0

replace github.com/adlerhurst/cli-client => ..
