package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"log"
	"makesite/packages/userutils"
	"strings"

	"os"
)
type Post struct {
	Content string
}

func main() {
	//Handles potential system arguments called on execution.
	fileFlag := flag.String("file", "html-file.html", "Tells the program what '.txt' file to convert.")
	dirFlag := flag.String("dir", "", "Specifies the directory in which to search for all '.txt' files.")
	
	flag.Parse()

	paths := []string {
		"template.tmpl",
	}

	if *dirFlag != "" {
		files, err := ioutil.ReadDir(*dirFlag)
		if err != nil {
			log.Fatal(err)
		}
		
		for _, file := range files {
			r,w,e := os.Pipe()
			if e != nil {
				panic(e)
			}

			os.Stdout = w
			if strings.Contains(file.Name(),".txt") {
				post := Post{Content:userutils.ReadFile(file.Name())}
				paths = append(paths, file.Name())
				t := template.Must(template.New("template.tmpl").ParseFiles(paths...))
				err = t.Execute(w, post)
				if err != nil {
					panic(err)
				}
				w.Close()
				out, e := ioutil.ReadAll(r)
				if e != nil {
					panic(e)
				}
				if *fileFlag == "" {
					userutils.WriteFile(file.Name()[0:len(file.Name())-4]+".html",out)
				} else {
					userutils.WriteFile(*fileFlag, out)
					
				}
			}
		}
	} else if *fileFlag != ""{
		r,w,e := os.Pipe()
			if e != nil {
				panic(e)
			}
		t := template.Must(template.New("template.tmpl").ParseFiles(paths...))
		err := t.Execute(w, Post{Content: userutils.ReadFile("first-post.txt")})
		if err != nil {
			panic(err)
		}

		w.Close()

		out, e := ioutil.ReadAll(r)
		if e != nil {
			panic(e)
		}

		userutils.WriteFile(*fileFlag, out)

	}

	
	

}