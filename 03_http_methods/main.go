package main

/*
	Etapa 3:
	Objetivo: Adicionar suporte a diferentes métodos HTTP (GET, POST, PUT, DELETE) para a rota "/".
	Isso simula o comportamento de uma API RESTful, onde cada método representa uma operação diferente.
	As outras rotas (/hello e /bye) continuam simples, apenas respondendo a requisições GET.
*/

import (
	"fmt"
	"net/http"
)

func main() {
	// Cada rota é registrada com um handler específico
	http.HandleFunc("/", home)       // Rota principal com múltiplos métodos HTTP
	http.HandleFunc("/hello", hello) // Rota simples: /hello
	http.HandleFunc("/bye", bye)     // Rota simples: /bye

	// Inicia o servidor HTTP na porta 8080
	http.ListenAndServe(":8080", nil)
}

// Handler da rota "/"
// Este handler verifica o método HTTP da requisição e responde de forma diferente para cada um
func home(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Welcome to the home page!")
	case "POST":
		fmt.Fprintf(w, "Post to the home page!")
	case "PUT":
		fmt.Fprintf(w, "Put to the home page!")
	case "DELETE":
		fmt.Fprintf(w, "Delete the home page!")
	default:
		// Retorna erro 405 caso o método HTTP não seja suportado
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
	}
}

// Handler da rota "/hello"
// Simples resposta a uma requisição GET
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

// Handler da rota "/bye"
// Simples resposta a uma requisição GET
func bye(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Goodbye guys!")
}
