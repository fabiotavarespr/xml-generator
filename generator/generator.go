package generator

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/fabiotavaresp/xml-generator/generator/config"
	"github.com/sirupsen/logrus"
)

// Generator struct owns items needed to start the service
type Generator struct {
	*config.GeneratorBuilder
}

// Run generator
func (g *Generator) Run(generatorBuilder *config.GeneratorBuilder) *Generator {
	logrus.Info("Starting XML Generator...")
	g.GeneratorBuilder = generatorBuilder

	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		fmt.Println(f.Name())
	}

	select {}
}
