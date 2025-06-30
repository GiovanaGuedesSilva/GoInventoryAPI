package main

/*
	Etapa 1 - Servidor HTTP Básico com net/http

	Esta primeira etapa mostra como criar uma API RESTful mínima em Go
	utilizando apenas a biblioteca padrão `net/http`.

	A aplicação escuta na porta 8080 e responde com um JSON contendo uma mensagem
	simples para qualquer requisição feita na rota raiz ("/").

	Esse é o ponto de partida para evoluir a aplicação em etapas mais estruturadas.
*/

import (
	"encoding/json" // Utilizado para codificar dados em formato JSON
	"net/http"      // Biblioteca padrão para criar servidores HTTP em Go
)

func main() {
	// Associa a função handler à rota "/"
	http.HandleFunc("/", handler)

	// Inicia o servidor na porta 8080
	http.ListenAndServe(":8080", nil)
}

// handler é a função responsável por lidar com requisições à rota "/"
func handler(w http.ResponseWriter, r *http.Request) {
	// Retorna um JSON com a mensagem "Hello World"
	json.NewEncoder(w).Encode(map[string]string{"message": "Hello World"})
}
