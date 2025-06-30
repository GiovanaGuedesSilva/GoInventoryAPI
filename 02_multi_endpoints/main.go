package main

/*
	Etapa 2:
	Objetivo: Adicionar novas rotas além da principal.
	Cada rota (/, /hello, /bye) tem sua própria função handler.
	Isso ilustra como a API pode crescer para ter múltiplos endpoints.
*/

import (
	"fmt"
	"net/http"
)

func main() {
	// Cada rota é registrada com um handler específico
	http.HandleFunc("/", home)       // Rota principal
	http.HandleFunc("/hello", hello) // Nova rota: /hello
	http.HandleFunc("/bye", bye)     // Nova rota: /bye

	// Inicia o servidor HTTP na porta 8080
	http.ListenAndServe(":8080", nil)
}

// Handler da rota "/"
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
}

// Handler da rota "/hello"
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

// Handler da rota "/bye"
func bye(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GoodBye guys!")
}
