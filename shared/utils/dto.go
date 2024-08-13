package utils

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

func UnmarshallDto[T any](data []byte) (T, error) {
	var res T
	err := json.Unmarshal(data, &res)
	if err != nil {
		logrus.Error(err.Error())
		return res, err
	}
	return res, nil
}
