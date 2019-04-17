// The following directive is necessary to make the package coherent:

// +build ignore

// This program generates insecure.gen.go. It can be invoked by running
// go generate
package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
	"time"
)

const envInsecure = "OCF_INSECURE"

// run with generateInsecure [package]
func main() {
	if len(os.Args) != 2 {
		die(fmt.Errorf("run with package argument: %v PACKAGE_NAME", os.Args[0]))
	}
	packageName := os.Args[1]

	f, err := os.Create("insecure.gen.go")
	die(err)
	defer f.Close()

	insecure := false
	insecureStr := os.Getenv(envInsecure)
	okayResponses := []string{"y", "Y", "yes", "Yes", "YES", "true", "TRUE"}
	for _, okayResponse := range okayResponses {
		if okayResponse == insecureStr {
			insecure = true
			break
		}
	}
	envInsecure := ""
	if len(insecureStr) > 0 {
		envInsecure = "envInsecure=" + insecureStr + " "
	}

	packageTemplate.Execute(f, struct {
		Timestamp   time.Time
		EnvInsecure string
		Package     string
		Insecure    bool
	}{
		Timestamp:   time.Now(),
		EnvInsecure: envInsecure,
		Package:     packageName,
		Insecure:    insecure,
	})
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var packageTemplate = template.Must(template.New("").Parse(`// Code generated by go generate; DO NOT EDIT.
// This file was generated via command "{{ .EnvInsecure }}go generate ./..." at
// {{ .Timestamp }}
package {{ .Package }}

// IsInsecure if returns true, listener and connections are without TLS
func IsInsecure() bool {
	return {{ .Insecure }}
}
`))
