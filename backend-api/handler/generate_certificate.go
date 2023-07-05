package handler

import (
	"github.com/achaquisse/skulla-api/internal"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type CertificateInput struct {
	Template    string `validate:"required"`
	StudentName string `validate:"required,min=5"`
	Description string `validate:"required,min=20"`
}
type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var validate = validator.New()

func validateStruct(input CertificateInput) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(input)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func GenerateCertificate(c *fiber.Ctx) error {
	input := new(CertificateInput)

	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errors := validateStruct(*input)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	bytes, err1 := internal.GenerateCertificate(input.Template, input.StudentName, input.Description)
	if err1 != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err1.Error(),
		})
	} else {
		c.Set("Content-Type", "application/pdf")
		_, err := c.Response().BodyWriter().Write(bytes)
		return err
	}
}
