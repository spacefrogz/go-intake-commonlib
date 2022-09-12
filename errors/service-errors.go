package errors

import (
	"encoding/json"
	"fmt"
	"os"

	"golang.org/x/exp/slices"
)

type IntakeError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func readJsonFile(zone string) ([]byte, error) {
	pwd, err := os.Getwd()

	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf(pwd+"/errors/%s-errors.json", zone)
	bytes, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func SE(code string) *IntakeError {

	unknownError := &IntakeError{Code: "void", Message: "Unknown error."}

	bytes, err := readJsonFile("service")

	if err != nil {
		return unknownError
	}

	var intakeErrors []IntakeError
	err = json.Unmarshal(bytes, &intakeErrors)

	if err != nil {
		return unknownError
	}

	idx := slices.IndexFunc(intakeErrors, func(intakeError IntakeError) bool {
		return intakeError.Code == code
	})

	if idx < 0 {
		return unknownError
	}

	return &intakeErrors[idx]
}
