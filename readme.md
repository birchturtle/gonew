# GoNew ReadMe

## Introduction
Very simple, quick 'n dirty cli utility which uses the `text/template` package for creating new go templates for quickstarting new projects etc.
It works running `gonew` along with any wanted extra settings and outputs a .go file set up completely with package and imports declarations and a main-func if package is main.

## Usage
Simplest usage is simply running `gonew` without any arguments, using only defaults. 

The defaults are:
- type=cons
- package=main
- db=none
- imports=fmt (for writing hello, world!)
- output=main.go

This will results in a file looking like so:
```
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, GoNew World!")
}
```
The different settings should be pretty obvious, and the available flags are as follows:
- "-o mynewgofile.go", sets the file name of the output file, defaults to "main.go".
- "-p data", sets the package, defaults to "main".
- "-db sqlite", adds a dependency to the specified database driver to imports section, defaults to none. 	
	- Available providers are:
		* "sqlite" = _"github.com/mattn/go-sqlite3"
		* "postgres" = _"github.com/ps...."
- "-t web", sets an overall "type" for the program with some more defaults added automatically, e.g. "web" adds import to "net/http" and sets up a minimal "hello world"-server in a main function (if pkg is main). Defaults to "cons" which imports "fmt". Currently only these 2 types are available

Any args added *after* any flags set (all flags are optional) should be importable packages, e.g. os, flag, crypto etc.

So, as a little more interesting example, I might want to set up a webserver quickly in a main package, call the file server.go and I know that I want it to use an sqlite database for data persistence and some flags. I would then run `gonew -t web -db sqlite -o server.go flag` which would result in a file called server.go in the current directory looking like so:
```
package main

import (
	"fmt"
	"net/http"
	"database/sql"
	"flag"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	net.ListenAndServe(":8080", func(w http.Responsewriter, r *http.Response) {
	fmt.Fprintf(w, "Hello, GoNew Web-World!")
}
```

just like that!

Also note, the imports are always "fmt" first, possibly followed by "net/http" and/or "database/sql", specific database driver at the bottom, and everything in between is added in the order you add them on the command line; You may want to run gofmt or equivalent on the new file. 
