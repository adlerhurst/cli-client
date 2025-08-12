package parse

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var (
	Path      string
	Json      string
	FromFlags bool
)

func InitFlags(cmd *cobra.Command, args []string) error {
	payloadSet := pflag.NewFlagSet("cli-request", pflag.ContinueOnError)
	payloadSet.BoolVar(&FromFlags, "payload", false, "using flags to set request payload")
	payloadSet.StringVar(&Path, "path", "", "path to json file containing request payload")
	payloadSet.StringVar(&Json, "json", "", "json object as params")
	payloadSet.ParseErrorsWhitelist.UnknownFlags = true
	cmd.Flags().AddFlagSet(payloadSet)

	cmd.MarkFlagFilename("path")
	cmd.MarkFlagsMutuallyExclusive("json", "path", "payload")
	cmd.MarkFlagsOneRequired("json", "path", "payload")

	return payloadSet.Parse(args)
}

func PreRun(request any) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		ParseArg(request, args)
	}
}

func ParseArg(request any, args []string) {
	if len(args) == 0 {
		fmt.Printf("no args provided")
		os.Exit(1)
	}
	if len(args) > 1 {
		fmt.Printf("too many args provided")
		os.Exit(1)
	}
	if json.Valid([]byte(args[0])) {
		if err := json.Unmarshal([]byte(args[0]), request); err != nil {
			fmt.Print(fmt.Sprintf("failed to unmarshal json arg: %v", err))
			os.Exit(1)
		}
		return
	}
	file, err := os.Open(args[0])
	if err != nil {
		fmt.Printf("failed to open file path: %v", err)
		os.Exit(1)
	}
	defer file.Close()
	if err := json.NewDecoder(file).Decode(request); err != nil {
		fmt.Print(fmt.Sprintf("failed to decode file file: %v", err))
		os.Exit(1)
	}
}
