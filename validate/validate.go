package validate

import (
	"encoding/json"
	"errors"
	"io"
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

type validate struct {
	v     *validator.Validate
	trans ut.Translator
}

func New() (*validate, error) {
	eng := en.New()
	uni := ut.New(eng, eng)

	t, found := uni.GetTranslator("en")
	if !found {
		return nil, errors.New("translator not found")
	}

	v := validator.New()

	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	if err := enTranslations.RegisterDefaultTranslations(v, t); err != nil {
		return nil, err
	}

	return &validate{v: v, trans: t}, nil
}

// Translate translates error into map of errors where key is field name and value is error.
func (v *validate) Translate(err error) map[string]string {
	errs := make(map[string]string)
	for _, err := range err.(validator.ValidationErrors) {
		errs[err.Field()] = err.Translate(v.trans)
	}
	return errs
}

// Struct is a helper method for validator.Validate.Struct.
func (v *validate) Struct(s any) error { return v.v.Struct(s) }

// Into decodes io.Reader impl to dest and validates it with Struct.
// Created for using mostly with http.Body, dest must be a pointer.
func (v *validate) Into(r io.Reader, dest any) error {
	if err := json.NewDecoder(r).Decode(dest); err != nil {
		return err
	}

	return v.Struct(dest)
}
