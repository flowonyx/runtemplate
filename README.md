`runtemplate` provides a way of doing template execution from standard Go templates from the command line.

It is intended primarily to be used with `go generate`. Simply put the `go generate` comment in your code like this:

`//go:generate runtemplate filename.tpl outfile.go Option1=Value1 Option2=Value2`

Then run `go generate` and it will run this against the specified template,
passing in whatever options have have been specified on the command line as a map.

In the template file, you can then access them simply by their keys. For instance:

`{{ .Option1 }}`

`.OutFile` and `.TemplateFile` are always available to the templates.

Also included are a couple of filters that may be helpful.

* title - Converts the input to Title Case.
* upper - Converts the input to UPPER CASE.
* lower - Converts the input to lower case.
* splitDotFirst - Given an input that has a '.' separator, returns the part before the first '.'.
* splitDotLast - Given an input that has a '.' separator, returns the part after the last '.'.

The last two are useful for getting only the package name or only the type name if passed an input of `package.Type`.
