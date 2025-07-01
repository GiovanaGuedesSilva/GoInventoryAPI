package config

import "errors"

/*
	ErrNotFound é um erro reutilizável que representa a situação
	em que um recurso (ex: item) não foi encontrado.

	Ele pode ser usado em qualquer lugar da aplicação (repositórios, usecases, handlers),
	permitindo que a camada superior (ex: handler HTTP) saiba que deve responder com 404 (Not Found).
*/
var ErrNotFound = errors.New("not found")
