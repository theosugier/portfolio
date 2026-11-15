package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
)

var _temp = template.Must(template.ParseGlob("./templates/*.html"))

type Project struct {
	Id          int
	Title       string
	Image       string
	Stack       []string
	Intro       string
	Description string
	Features    []string
	Repo        string
	Video       string
}

var listProjects []Project
