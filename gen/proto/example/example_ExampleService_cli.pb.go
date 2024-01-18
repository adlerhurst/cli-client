// Code generated by protoc-gen-go-cli. DO NOT EDIT.

package example

import (
	cobra "github.com/spf13/cobra"
	pflag "github.com/spf13/pflag"
)

var (
	ExampleCmd = &cobra.Command{
		Use: "example",
		Short: ` This is an example service
 The service does nothing
`,
		Long: ` This is an example service
 The service does nothing
`,
		FParseErrWhitelist: cobra.FParseErrWhitelist{UnknownFlags: true},
		DisableFlagParsing: true,
	}
)

var (
	_ExampleMyCallCmdRequest = &CallRequestFlag{CallRequest: new(CallRequest)}
	ExampleMyCallCmd         = &cobra.Command{
		Use: "mycall",
		Short: ` I do absolutely nothing
`,
		Long: ` I do absolutely nothing
`,
		Run:                runExampleMyCallCmd,
		FParseErrWhitelist: cobra.FParseErrWhitelist{UnknownFlags: true},
		DisableFlagParsing: true,
	}
)

func init() {
	ExampleCmd.AddCommand(ExampleMyCallCmd)
	_ExampleMyCallCmdRequest.AddFlags(ExampleMyCallCmd.Flags())
	ExampleMyCallCmd.PreRun = func(cmd *cobra.Command, args []string) {
		_ExampleMyCallCmdRequest.ParseFlags(cmd.Flags(), args)
	}
}

func runExampleMyCallCmd(cmd *cobra.Command, args []string) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		if !f.Changed {
			return
		}
		DefaultConfig.Logger.Info("flag", "name", f.Name, "value", f.Value.String(), "changed", f.Changed)
	})
}

var (
	_ExampleMyNestedCallCmdRequest = &NestedRequestFlag{NestedRequest: new(NestedRequest)}
	ExampleMyNestedCallCmd         = &cobra.Command{
		Use:                "mynestedcall",
		Short:              ``,
		Long:               ``,
		Run:                runExampleMyNestedCallCmd,
		FParseErrWhitelist: cobra.FParseErrWhitelist{UnknownFlags: true},
		DisableFlagParsing: true,
	}
)

func init() {
	ExampleCmd.AddCommand(ExampleMyNestedCallCmd)
	_ExampleMyNestedCallCmdRequest.AddFlags(ExampleMyNestedCallCmd.Flags())
	ExampleMyNestedCallCmd.PreRun = func(cmd *cobra.Command, args []string) {
		_ExampleMyNestedCallCmdRequest.ParseFlags(cmd.Flags(), args)
	}
}

func runExampleMyNestedCallCmd(cmd *cobra.Command, args []string) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		if !f.Changed {
			return
		}
		DefaultConfig.Logger.Info("flag", "name", f.Name, "value", f.Value.String(), "changed", f.Changed)
	})
}
