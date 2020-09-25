package generator

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/fabiotavarespr/xml-generator/generator/config"
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

	listNames := fileNames(g.Source)
	createFiles(g.Destination, listNames)

	return g
}

func fileNames(source string) []string {
	listNames := make([]string, 0)

	files, err := ioutil.ReadDir(source)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		listNames = append(listNames, f.Name())
	}

	return listNames

}

func createFiles(destination string, listNames []string) {
	for _, f := range listNames {
		file, err := os.Create(destination + strings.Replace(f, ".png", ".xml", -1))
		defer file.Close()

		if err != nil {
			log.Fatal(err)
		}
		_, err2 := file.WriteString("<img src=\"" + f + "\"></img>")

		if err2 != nil {
			log.Fatal(err2)
		}
	}
}
