package main

import (
	"flag"
	"os"
	"text/template"
)

// OFile for output file struct
type OFile struct {
	Pkg      string   // -p, default main
	Db       string   // -db, default ""
	Imports  []string // remaining args
	Type     string   // default "cons", options "cons", "web"
	Filename string   // -o, default "main.go"
}

var databases = make(map[string]string)

var content string = `
package {{ .Pkg }}

import (
{{ if eq .Pkg "main" }}
	{{"\"fmt\"" -}}
{{ end }}
{{ if eq .Type "web"}}
	{{ "\"net/http\"" -}}
{{ end }}
{{ range .Imports }}
	"{{ . -}}"
{{ end }}

{{ if ne .Db  ""}}
	{{ .Db }}
{{ end }}
)
{{ if (eq .Type "web")}}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, GoNew Web-World!") 
}
{{ end }}
{{ if (eq .Pkg "main")}}
func main() {
	{{ if eq .Type "cons" }}
	fmt.Println("Hello, GoNew World!")
	{{ else if eq .Type "web" }}
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
	{{ end }}
}
{{ end }}
`

func init() {
	databases["sqlite"] = `_"github.com/mattn/go-sqlite3"`
	databases[""] = ""
}

func main() {
	of := OFile{}
	var t string
	db := flag.String("db", "", "set to nothing if not set")
	flag.StringVar(&of.Filename, "o", "main.go", "= main.go if not set")
	flag.StringVar(&of.Pkg, "p", "main", "= main if not set")
	flag.StringVar(&t, "t", "cons", "type of application, cons if not set; Vailable: web, cons")
	flag.Parse()
	setDb, ok := databases[*db]
	if ok {
		of.Db = setDb
	} else {
		panic("err: unknown type db - available: sqlite, postgres")
	}
	var imports []string
	for _, arg := range flag.Args() {
		imports = append(imports, arg)
	}
	of.Imports = imports
	of.Type = t
	templ, err := template.New("test").Parse(content)
	if err != nil {
		panic(err)
	}
	output, err := os.Create(of.Filename)
	if err != nil {
		panic(err)
	}
	err = templ.Execute(output, of)
	if err != nil {
		panic(err)
	}
}
