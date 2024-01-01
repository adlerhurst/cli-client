// Code generated by protoc-gen-go-cli. DO NOT EDIT.
//23:38:50

package example

import (
	cobra "github.com/spf13/cobra"
	pflag "github.com/spf13/pflag"
)

var (
	_ExampleMyNestedCallCmdRequest = new(NestedRequest)
	ExampleMyNestedCallCmd         = &cobra.Command{
		Use:   "mynestedcall",
		Short: ``,
		Long:  ``,

		PreRun:             _ExampleMyNestedCallCmdRequest.ParseFlags,
		Run:                runExampleMyNestedCallCmd,
		FParseErrWhitelist: cobra.FParseErrWhitelist{UnknownFlags: true},
		DisableFlagParsing: true,
	}
)

func runExampleMyNestedCallCmd(cmd *cobra.Command, args []string) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		DefaultConfig.Logger.Info("flag", "name", f.Name, "value", f.Value.String(), "changed", f.Changed)
	})
}
