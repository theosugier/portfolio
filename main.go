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
