package dto_utils

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func CustomBodyValidator(content []byte, target interface{}) (uint, error) {
	// TODO : protect against very big bodys !
	if err := json.NewDecoder(bytes.NewReader(content)).Decode(&target); err != nil {
		return http.StatusBadRequest, err
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(target)

	if err != nil {
		return http.StatusBadRequest, err
	}

	return 0, nil
}
