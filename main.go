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

	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		log.Fatalf("Erreur (Serveur) : Impossible de lancer le serveur : %v", err)
		return
	}
}

func renderTemplate(w http.ResponseWriter, name string, data interface{}) {
	buffer := bytes.Buffer{}

	errRender := _temp.ExecuteTemplate(&buffer, name, data)
	if errRender != nil {
		log.Printf("Erreur (Template) : Impossible de rendre le template %s : %v", name, errRender)
		http.Error(w, "Erreur chargement de la page", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	buffer.WriteTo(w)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request : %s %s", r.Method, r.URL.Path)

	renderTemplate(w, "index", nil)
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

func ProjectsHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request : %s %s", r.Method, r.URL.Path)

	renderTemplate(w, "projects-list", listProjects)
}

func ProjectByIdHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request : %s %s", r.Method, r.URL.Path)

	id, errId := strconv.Atoi(r.PathValue("id"))
	if errId != nil {
		http.Error(w, "Erreur données invalides...", http.StatusBadRequest)
		return
	}

	for _, project := range listProjects {
		if project.Id == id {
			renderTemplate(w, "project-details", project)
			return
		}
	}
	http.Error(w, "Projet introuvable...", http.StatusNotFound)
}

func NotfoundHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request : %s %s", r.Method, r.URL.Path)

	renderTemplate(w, "error", nil)
}
