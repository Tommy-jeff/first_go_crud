package resterr

import "net/http"

/// Neste arquivos definimos o objeto que será retornado em qualquer erro que nossa aplicação posso lançar

/// Esse é o objeto base que será retornado para erros
/// Os campos `json:` tem a finalidade de direcionar como a conversão dele para uma estrutura json ficará
type RestErr struct {
	Message string 	 `json:"message"`
	Code  	int    	 `json:"code"`
	Err     string 	 `json:"error"`
	Causes 	[]Causes `json:"causes,omitempty"`
}

/// Objeto de causas que compõe o objeto `RestErr`
type Causes struct {
	Field 	string `json:"field"`
	Message string `json:"message"`
}

/// Metodo que retorna a mensagem do erro
func(r *RestErr) Error() string {
	return r.Message
}

/// Metodo que cria o objeto `RestErr` padrão
func NewRestErr(message string, code int, err string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Code:    code,
		Err:   	 err,
		Causes:  causes,
	}
}

/// Metodo que cria o objeto `RestErr` do tipo `BadRequestError`
func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Code:    http.StatusBadRequest,
		Err:     "bad_request",
	}
}

/// Metodo que cria o objeto `RestErr` do tipo `BadRequestValidationError`
func NewBadRequestValidationError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Code:    http.StatusBadRequest,
		Err:   	 "bad_request",
		Causes:  causes,
	}
}

/// Metodo que cria o objeto `RestErr` do tipo `InternalServerError`
func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Code:    http.StatusInternalServerError,
		Err:     "internal_server_error",
	}
}

/// Metodo que cria o objeto `RestErr` do tipo `NotFoundError`
func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Code:    http.StatusNotFound,
		Err:     "not_found",
	}
}

/// Metodo que cria o objeto `RestErr` do tipo `ForbiddenError`
func NewForbiddenError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Code:    http.StatusForbidden,
		Err:     "forbidden",
	}
}   