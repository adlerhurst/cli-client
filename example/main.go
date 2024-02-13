package main

import (
	"github.com/adlerhurst/cli-client/example/cli"
	"github.com/spf13/cobra"
)

func main() {
	err := cli.ExampleCmd.Execute()
	cobra.CheckErr(err)
}
