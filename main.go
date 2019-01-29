package main

import (
	"gopkg.in/go-playground/colors.v1"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func randomColor() colors.Color {
	s := rand.NewSource(time.Now().UnixNano())
	random := rand.New(s)

	r := uint8(random.Intn(256))
	g := uint8(random.Intn(256))
	b := uint8(random.Intn(256))

	color, _ := colors.RGB(r, g, b)

	return color
}

func main() {
	color := randomColor()

	tmpl := template.Must(template.ParseFiles("index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_ = tmpl.Execute(w, template.CSS(color.String()))
	})

	log.Fatal(http.ListenAndServe(":80", nil));
}

