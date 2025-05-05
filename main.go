package main

import (
    "html/template"
    "io/ioutil"
    "log"
    "net/http"
    "path/filepath"
    "strings"
)

type Post struct {
    Title   string
    Content string
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
    http.HandleFunc("/", indexHandler)

    http.HandleFunc("/post/", postHandler)

    log.Println("Starting server on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
    files, err := filepath.Glob("posts/*.md")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    var posts []Post
    for _, file := range files {
        content, err := ioutil.ReadFile(file)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        lines := strings.SplitN(string(content), "\n", 2)
        posts = append(posts, Post{Title: lines[0], Content: lines[1]})
    }

    templates.ExecuteTemplate(w, "index.html", posts)
}


func postHandler(w http.ResponseWriter, r *http.Request) {

    postName := strings.TrimPrefix(r.URL.Path, "/post/") + ".md"
    content, err := ioutil.ReadFile("posts/" + postName)
    if err != nil {
        http.NotFound(w, r)
        return
    }

    lines := strings.SplitN(string(content), "\n", 2)
    post := Post{Title: lines[0], Content: lines[1]}

    templates.ExecuteTemplate(w, "post.html", post)
}
