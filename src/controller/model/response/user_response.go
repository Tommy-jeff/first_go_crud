package response

/// Um modelo de response que deve ser retornado
type UserResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int8 	 `json:"age"`
}