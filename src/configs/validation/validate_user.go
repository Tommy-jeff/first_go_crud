package validation

import (
	"encoding/json"
	"errors"

	resterr "github.com/Tommy-jeff/first_go_crud/src/configs/rest_err"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var transl ut.Translator

/// Inicializa o translator, sua chamada é automática pelo runtime sempre que o package (validation) é importado, ou seja, não precisa ser invocada manualmente
/// Importante: o valor passado como parâmetro (*validator.Validate) vem direto do gin-gonic, isso porque ele usa o pacote validator/v10 e já o inicializa internamente 
func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
} 

/// Função que retorna o erro de validação de um user
/// internamente ele compara o erro inserido com os erros de validação do pacote validator/v10 e retorna um ubjeto de `RestErr`
func ValidateUserError(validation_err error) *resterr.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return resterr.NewBadRequestError("Invalid Field Type")
	} else if errors.As(validation_err, &jsonValidationError) {
		errorsCauses := []resterr.Causes{}

		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := resterr.Causes{
				Message: e.Translate(transl),
				Field: e.Field(),
			}

			errorsCauses = append(errorsCauses, cause)
		}

		return resterr.NewBadRequestValidationError("Some fields are invalid", errorsCauses)
	} else {
		return resterr.NewBadRequestError("Error when trying to validate user")
	}
}