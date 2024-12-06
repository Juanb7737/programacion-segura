package main

import (
    "fmt"
    "net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Hola Mundo</h1>")
}

func main() {
	// Servir archivos estáticos desde el directorio "static"
    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/", fs)
	
    http.HandleFunc("/api/hello", home)
    
    fmt.Println("Servidor iniciado en http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}