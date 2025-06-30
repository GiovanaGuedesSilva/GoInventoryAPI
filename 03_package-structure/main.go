package main

/*
	Etapa 3 - Suporte a Múltiplos Métodos HTTP

	Nesta etapa, o handler é adaptado para diferenciar o tratamento
	conforme o método HTTP usado (GET, POST, etc).

	Esse controle é essencial em APIs RESTful para permitir operações diferentes
	com base no método (ex: GET para leitura, POST para criação).
*/

import (
	"encoding/json"
	"net/http"
)

func main() {
	// Define o handler para a rota "/"
	http.HandleFunc("/", handler)

	// Inicia o servidor na porta 8080
	http.ListenAndServe(":8080", nil)
}

// handler lida com diferentes métodos HTTP
func handler(w http.ResponseWriter, r *http.Request) {
	// Analisa o método da requisição
	switch r.Method {
	case http.MethodGet:
		// Responde com mensagem para requisição GET
		json.NewEncoder(w).Encode(map[string]string{"message": "GET request received"})

	case http.MethodPost:
		// Responde com mensagem para requisição POST
		json.NewEncoder(w).Encode(map[string]string{"message": "POST request received"})

	default:
		// Retorna erro para métodos não suportados
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
