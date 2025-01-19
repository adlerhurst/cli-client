package cli

import (
	"reflect"
	"testing"

	"github.com/spf13/pflag"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type addHuman struct {
	Username []string
	Profile  struct {
		FirstName string
		LastName  string
	}
	Emails []struct {
		Address  string
		Verified bool
	}
	Type enum
}

type enum int32

// Descriptor implements protoreflect.Enum.
func (e enum) Descriptor() protoreflect.EnumDescriptor {
	panic("unimplemented")
}

// Number implements protoreflect.Enum.
func (e enum) Number() protoreflect.EnumNumber {
	panic("unimplemented")
}

// Type implements protoreflect.Enum.
func (e enum) Type() protoreflect.EnumType {
	panic("unimplemented")
}

var _ protoreflect.Enum = enum(0)

func Test(t *testing.T) {
	var human addHuman
	args := []string{
		"--username", "adlerhurst",
		"--username", "hurst",
		"--type", "ADMIN",
		"--profile", "--first-name", "Adler", "--last-name", "Hurst",
		"--email", "--address", "adler@hurst.dev", "--verified",
		"--email", "--address", "adler@hurst.com", "--verified=false",
	}

	humanFlags := pflag.NewFlagSet("human", pflag.ContinueOnError)
	newSliceValue(&human.Username).Add(humanFlags, "username", "", "")
	newSliceValue(&human.Type).Add(humanFlags, "type", "", "")

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
		endIdx:   5,
		sub: []indexTree{
			{
				flags:    profileFlags,
				startIdx: 7,
				endIdx:   10,
			},
			{
				flags:    emailFlags[0],
				startIdx: 12,
				endIdx:   14,
			},
			{
				flags:    emailFlags[1],
				startIdx: 16,
				endIdx:   18,
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
