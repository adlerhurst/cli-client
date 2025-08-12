package main

import (
	cli "github.com/adlerhurst/cli-client/example/gen/adlerhurst/example/v1"
	"github.com/spf13/cobra"
)

func main() {
	err := cli.ExampleServiceCommand.Execute()
	cobra.CheckErr(err)
}
