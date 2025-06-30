package main

/*
	Etapa 2 - Separação da Função Handler

	Nesta etapa, a principal melhoria é a separação da função handler
	do corpo da função `main`, tornando o código mais organizado, limpo e reutilizável.

	Essa separação facilita a manutenção, testes unitários e futuras extensões do handler.
*/

import (
	"encoding/json"
	"net/http"
)

func main() {
	// Define o handler da rota raiz
	http.HandleFunc("/", handler)

	// Inicia o servidor na porta 8080
	http.ListenAndServe(":8080", nil)
}

// handler é a função que trata as requisições HTTP para a rota "/"
func handler(w http.ResponseWriter, r *http.Request) {
	// Envia uma resposta em JSON com a mensagem padrão
	json.NewEncoder(w).Encode(map[string]string{"message": "Hello World"})
}
