package main

/*
	Etapa 3 - Suporte a Múltiplos Métodos HTTP

	Nesta etapa, estendemos a rota "/" para lidar com diferentes métodos HTTP (GET, POST, PUT, DELETE).
	Isso é essencial para construir uma API RESTful, em que cada método representa uma ação distinta sobre um recurso.

	As outras rotas ("/hello" e "/bye") seguem utilizando apenas o método GET, como nas etapas anteriores.
*/

import (
	"fmt"
	"net/http"
)

func main() {
	// Associa cada rota a sua respectiva função handler
	http.HandleFunc("/", home)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/bye", bye)

	// Inicia o servidor na porta 8080
	http.ListenAndServe(":8080", nil)
}

// home trata a rota "/" e suporta múltiplos métodos HTTP
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
		// Retorna erro se o método não for suportado
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
	}
}

// hello trata a rota "/hello" com uma resposta simples
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

// bye trata a rota "/bye" com uma mensagem de despedida
func bye(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Goodbye guys!")
}
