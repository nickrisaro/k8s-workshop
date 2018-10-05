package main

import (
	"fmt"
	"net/http"
)

func playHome(w http.ResponseWriter, r *http.Request) {
	html := `<!DOCTYPE html><html><body><center>
		<img src="https://raw.githubusercontent.com/twogg-git/k8s-workshop/master/src/1.0.png">
		<h1 style="color:blue">Practicando con Kubernetes</h1>
		<h3 style="color:blue">Versión: twogghub/web:1.0</h3>	
		</center></body></html>`
	fmt.Fprintf(w, html)
}

func main() {
	http.HandleFunc("/", playHome)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
