package main

/*
	Etapa 2 - Múltiplos Endpoints com net/http

	Nesta etapa, adicionamos múltiplas rotas ao servidor HTTP, cada uma com sua própria função handler.
	Isso permite estruturar melhor as funcionalidades da API, separando as responsabilidades por endpoint.

	As rotas implementadas são:
	- "/"      → Página inicial
	- "/hello" → Saudação
	- "/bye"   → Despedida
*/

import (
	"fmt"
	"net/http"
)

func main() {
	// Registra múltiplos handlers para rotas diferentes
	http.HandleFunc("/", home)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/bye", bye)

	// Inicia o servidor na porta 8080
	http.ListenAndServe(":8080", nil)
}

// home trata a rota raiz "/"
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
}

// hello trata a rota "/hello"
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

// bye trata a rota "/bye"
func bye(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GoodBye guys!")
}
