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

func main() {
	data, errFile := os.ReadFile("./data/projects.json")
	if errFile != nil {
		log.Fatalf("Erreur (JSON) : Impossible de lire le fichier : %v", errFile)
		return
	}
	errJson := json.Unmarshal(data, &listProjects)
	if errJson != nil {
		log.Fatalf("Erreur (JSON) : Impossible de lire les projets : %v", errJson)
		return
	}

	router := http.NewServeMux()

	router.HandleFunc("/{$}", HomeHandler)
	router.HandleFunc("/projects", ProjectsHandler)
	router.HandleFunc("/projects/{id}", ProjectByIdHandler)
	router.HandleFunc("/", NotfoundHandler)

	fileServer := http.FileServer(http.Dir("./assets"))
	router.Handle("/static/", http.StripPrefix("/static/", fileServer))

}

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
