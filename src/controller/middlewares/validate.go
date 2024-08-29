package middlewares

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/eron97/bff-golang.git/cmd/config/exceptions"
	"github.com/eron97/bff-golang.git/src/controller/dtos"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	en := en.New()
	unt := ut.New(en, en)
	transl, _ = unt.GetTranslator("en")
	en_translation.RegisterDefaultTranslations(Validate, transl)
}

func UserValidationMiddleware(ctx *fiber.Ctx) error {
	zap.L().Info("Starting user validation")

	var createUser dtos.CreateUser
	data := ctx.Body()

	err := ValidateUnexpectedFields(ctx, data)
	if err != nil {
		zap.L().Error("Unexpected fields in the request", zap.Error(err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := json.Unmarshal(data, &createUser); err != nil {
		zap.L().Error("Error when unmarshalling data", zap.Error(err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid field type",
		})
	}

	if err := Validate.Struct(&createUser); err != nil {
		var jsonValidationError validator.ValidationErrors
		if errors.As(err, &jsonValidationError) {
			errorsCauses := []exceptions.Causes{}
			for _, e := range jsonValidationError {
				cause := exceptions.Causes{
					FieldMessage: e.Translate(transl),
					Field:        e.Field(),
				}
				errorsCauses = append(errorsCauses, cause)
			}
			zap.L().Error("Error validating fields", zap.Error(err))
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"request invalid": exceptions.NewBadRequestValidationError("Some fields are invalid", errorsCauses),
			})
		}

		zap.L().Info("Error converting fields", zap.Error(err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error trying to convert fields",
		})
	}

	ctx.Locals("createUser", createUser)
	zap.L().Info("User validation completed successfully", zap.Error(err))
	return ctx.Next()
}

func ValidateUnexpectedFields(ctx *fiber.Ctx, data []byte) error {

	zap.L().Info("Validating unexpected fields")

	var rawMap map[string]interface{}

	if err := json.Unmarshal(data, &rawMap); err != nil {
		zap.L().Error("Formato de JSON inválido", zap.Error(err))
		return exceptions.NewBadRequestError("Invalid JSON format")
	}

	expectedFields := map[string]bool{
		"first_name": true,
		"last_name":  true,
		"email":      true,
		"cep":        true,
		"country":    true,
		"city":       true,
		"address":    true,
		"password":   true,
	}

	var unexpectedFields []string
	for field := range rawMap {
		if !expectedFields[field] {
			unexpectedFields = append(unexpectedFields, field)
		}
	}

	if len(unexpectedFields) == 0 {
		return nil
	}

	zap.L().Info("Validating unexpected fields")
	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error": fmt.Sprintf("Unexpected fields: %v. Please remove them and try again.", unexpectedFields),
	})

}
