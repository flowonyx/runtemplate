// runtemplate provides a simple way to execute a template from the command line.
// It is called like this:
// runtemplate template.tpl out.go key1=value1 key2=value2
//
// This will execute the templated found in "template.tpl", passing it a map with the structure:
//   map[string]string{"OutFile": "out.go", "TemplateFile": "template.tpl", "key1": "value1", "key2": "value2"}
//
// The result will be written to the "out.go" file.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

func main() {

	// Context will be passed to the template as a map.
	context := make(map[string]string)

	args := os.Args[1:]

	templatePath := args[0]
	outPath := args[1]

	// set up some special context values just in case they are wanted.
	context["OutFile"] = outPath
	context["TemplateFile"] = templatePath

	// args cuts off the first 3 args because they are the name of the executable,
	// the template path, and the output path. We only want the key=value pairs in the context.
	for _, arg := range args[2:] {
		if !strings.Contains(arg, "=") {
			fmt.Printf("%s is not a valid argument. Should be in the format of key=value.", arg)
			os.Exit(1)
		}
		keyVal := strings.Split(arg, "=")
		context[keyVal[0]] = keyVal[1]
	}

	// read the template and if it is not there or there is some
	// other error, print the error and return
	b, err := ioutil.ReadFile(templatePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Set up some text munging functions that will be available in the templates.
	funcMap := template.FuncMap{
		"title": strings.Title,
		"lower": strings.ToLower,
		"upper": strings.ToUpper,
		// splitDotFirst returns the first part of a string split on a "."
		// Useful for the case in which you want the package name from a passed value
		// like "package.Type"
		"splitDotFirst": func(s string) string {
			parts := strings.Split(s, ".")
			return parts[0]
		},
		// splitDotLast returns the last part of a string split on a "."
		// Useful for the case in which you want the type name from a passed value
		// like "package.Type"
		"splitDotLast": func(s string) string {
			parts := strings.Split(s, ".")
			return parts[len(parts)-1]
		},
	}

	tmpl, err := template.New(templatePath).Funcs(funcMap).Parse(string(b))

	f, err := os.Create(outPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	err = tmpl.Execute(f, context)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Completed generating: ", outPath)

}
