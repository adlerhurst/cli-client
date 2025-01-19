package cli

import (
	"reflect"
	"testing"

	cli_client "github.com/adlerhurst/cli-client"
	cli "github.com/adlerhurst/cli-client/example/cli"
	"github.com/spf13/pflag"
	"google.golang.org/protobuf/types/known/anypb"
)

type addHuman struct {
	Username []string
	// Type     []cli.CallRequest_Wat
	Type    cli.CallRequest_Wat
	ts      *anypb.Any
	Profile struct {
		FirstName string
		LastName  string
	}
	Emails []struct {
		Address  string
		Verified bool
	}
}

func TestParse(t *testing.T) {
	req := new(cli.CallRequest)
	cli_client.ParseFlags(req)
}

func Test(t *testing.T) {
	var human addHuman
	args := []string{
		"--username", "adlerhurst",
		"--username", "hurst",
		"--type", "WAT_WEISS",
		"--ts", `{"@type": "google.protobuf.Struct", "value": {"key":"value"}}`,
		"--profile", "--first-name", "Adler", "--last-name", "Hurst",
		"--email", "--address", "adler@hurst.dev", "--verified",
		"--email", "--address", "adler@hurst.com", "--verified=false",
	}

	humanFlags := pflag.NewFlagSet("human", pflag.ContinueOnError)
	newSliceValue(&human.Username).Add(humanFlags, "username", "", "")
	newValue(&human.Type).Add(humanFlags, "type", "", "")
	newValue(&human.ts).Add(humanFlags, "ts", "", "")
	// newSliceValue(&human.Type).Add(humanFlags, "type", "", "")

	profileFlags := pflag.NewFlagSet("profile", pflag.ContinueOnError)
	newValue(&human.Profile.FirstName).Add(profileFlags, "first-name", "", "")
	newValue(&human.Profile.LastName).Add(profileFlags, "last-name", "", "")

	emailFlags := []*pflag.FlagSet{
		pflag.NewFlagSet("email", pflag.ContinueOnError),
		pflag.NewFlagSet("email", pflag.ContinueOnError),
	}
	for i, flagSet := range emailFlags {
		human.Emails = append(human.Emails, struct {
			Address  string
			Verified bool
		}{})
		newValue(&human.Emails[i].Address).Add(flagSet, "address", "", "")
		newValue(&human.Emails[i].Verified).Add(flagSet, "verified", "", "")
	}

	indexes := indexTree{
		flags:    humanFlags,
		startIdx: 0,
		endIdx:   7,
		sub: []indexTree{
			{
				flags:    profileFlags,
				startIdx: 9,
				endIdx:   12,
			},
			{
				flags:    emailFlags[0],
				startIdx: 14,
				endIdx:   16,
			},
			{
				flags:    emailFlags[1],
				startIdx: 18,
				endIdx:   20,
			},
		},
	}

	if err := indexes.parse(args); err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(
		human,
		addHuman{
			Username: []string{"adlerhurst", "hurst"},
			Type:     cli.CallRequest_WAT_ICH,
			// Type: []cli.CallRequest_Wat{cli.CallRequest_WAT_WEISS, cli.CallRequest_WAT_ICH},
			Profile: struct {
				FirstName string
				LastName  string
			}{
				FirstName: "Adler",
				LastName:  "Hurst",
			},
			Emails: []struct {
				Address  string
				Verified bool
			}{
				{
					Address:  "adler@hurst.dev",
					Verified: true,
				},
				{
					Address:  "adler@hurst.com",
					Verified: false,
				},
			},
		},
	) {
		t.Error("unexpected human:", human)
	}
}

func BenchmarkParse(b *testing.B) {
	var human addHuman
	args := []string{
		"--username", "adlerhurst",
		"--username", "hurst",
		"--profile", "--first-name", "Adler", "--last-name", "Hurst",
		"--email", "--address", "adler@hurst.dev", "--verified",
		"--email", "--address", "adler@hurst.com", "--verified=false",
	}

	humanFlags := pflag.NewFlagSet("human", pflag.ContinueOnError)
	newSliceValue(&human.Username).Add(humanFlags, "username", "", "")

	profileFlags := pflag.NewFlagSet("profile", pflag.ContinueOnError)
	newValue(&human.Profile.FirstName).Add(profileFlags, "first-name", "", "")
	newValue(&human.Profile.LastName).Add(profileFlags, "last-name", "", "")

	emailFlags := []*pflag.FlagSet{
		pflag.NewFlagSet("email", pflag.ContinueOnError),
		pflag.NewFlagSet("email", pflag.ContinueOnError),
	}
	human.Emails = make([]struct {
		Address  string
		Verified bool
	}, 2)
	for i, flagSet := range emailFlags {
		newValue(&human.Emails[i].Address).Add(flagSet, "address", "", "")
		newValue(&human.Emails[i].Verified).Add(flagSet, "verified", "", "")
	}

	indexes := indexTree{
		flags:    humanFlags,
		startIdx: 0,
		endIdx:   3,
		sub: []indexTree{
			{
				flags:    profileFlags,
				startIdx: 5,
				endIdx:   8,
			},
			{
				flags:    emailFlags[0],
				startIdx: 10,
				endIdx:   12,
			},
			{
				flags:    emailFlags[1],
				startIdx: 14,
				endIdx:   16,
			},
		},
	}

	b.Run("parse", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if err := indexes.parse(args); err != nil {
				b.Error(err)
			}
		}
	})

}
