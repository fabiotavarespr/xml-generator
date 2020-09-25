package cmd

import (
	"fmt"

	"github.com/fabiotavarespr/xml-generator/generator"
	"github.com/fabiotavarespr/xml-generator/generator/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var generatorCmd = &cobra.Command{

	Use:     "generator",
	Aliases: []string{"g"},
	Short:   "Starting XML Generator",
	Long:    "Starting XML Generator",
	RunE: func(cmd *cobra.Command, args []string) error {
		builder := new(config.GeneratorBuilder).Init(viper.GetViper())
		generator := new(generator.Generator).Run(builder)
		fmt.Println(generator)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(generatorCmd)

	config.AddFlags(generatorCmd.Flags())

	err := viper.GetViper().BindPFlags(generatorCmd.Flags())
	if err != nil {
		panic(err)
	}
}
