package request

/// Um modelo de request que deve ser recebido.
/// Os modelos json são declarodos juntos com "binding" o qual é usado para definir validações no request 
type UserRequest struct {
	Name 	 string `json:"name" binding:"required,min=4,max=100"`
	Email 	 string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,containsany=!@#$%&*"`
	Age 	 int8 	`json:"age" binding:"required,min=1,max=125"`
}