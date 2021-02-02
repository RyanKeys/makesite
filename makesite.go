package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"os"
)



type Post struct {
	Content string
}

type Posts struct {
	Items []Post
}

func main() {

	
	examplePtr := flag.String("file", "html-file.html", " Help text.")
	flag.Parse()
	
	// Assigns all template engines. AKA all user templates.
	paths := []string {
		"template.tmpl",
	}


	rescueStdout := os.Stdout
	r,w,_ := os.Pipe()
	os.Stdout = w
	
	//Creates 'posts' variable that contains a list of all text files contained in Post struct(s) instantiated above.
	posts := Posts {[]Post{{ReadFile("first-post.txt")},{ReadFile("latest-post.txt")}}}
	t := template.Must(template.New("template.tmpl").ParseFiles(paths...))
	err := t.Execute(os.Stdout, posts)
	if err != nil {
		panic(err)
	}
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	WriteFile(*examplePtr, out)
	

}


func ReadFile(filename string) string {
	fileContents, err := ioutil.ReadFile(filename)
        if err != nil {
            // A common use of `panic` is to abort if a function returns an error
            // value that we don’t know how to (or want to) handle. This example
            // panics if we get an unexpected error when creating a new file.
            panic(err)
        }
        return string(fileContents)
}

func WriteFile(filename string, data []byte) {
	err := ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		panic(err)
	}
}



