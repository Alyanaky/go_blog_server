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

