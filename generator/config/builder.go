package config

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	source      = "source"
	destination = "destination"
)

// Flags define the fields that will be passed via cmd
type Flags struct {
	Source      string
	Destination string
}

// GeneratorBuilder defines the parametric information of a server instance
type GeneratorBuilder struct {
	*Flags
}

// AddFlags adds flags for Builder.
func AddFlags(flags *pflag.FlagSet) {
	flags.StringP(source, "s", "", "Source directory")
	flags.StringP(destination, "d", "", "Destination directory")
}

// Init initializes the web server builder with properties retrieved from Viper.
func (b *GeneratorBuilder) Init(v *viper.Viper) *GeneratorBuilder {
	flags := new(Flags)
	flags.Source = v.GetString(source)
	flags.Destination = v.GetString(destination)
	flags.check()

	b.Flags = flags

	return b
}

func (flags *Flags) check() {
	logrus.Infof("Flags: '%v'", flags)

	requiredFlags := []struct {
		value string
		name  string
	}{
		{flags.Source, source},
		{flags.Destination, destination},
	}

	var errMsg string

	for _, flag := range requiredFlags {
		if flag.value == "" {
			errMsg += fmt.Sprintf("\n\t%v", flag.name)
		}
	}

	if errMsg != "" {
		errMsg = "The following flags are missing: " + errMsg
		panic(errMsg)
	}
}
