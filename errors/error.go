package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"golang.org/x/exp/slices"
)

func UnknownError() *IntakeError {
	return &IntakeError{Code: "void", Message: "Unknown error."}
}

type IntakeError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func readJsonFile(zone string) ([]byte, error) {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))

	path := fmt.Sprintf(filepath.Dir(d)+"/errors/%s-errors.json", zone)
	bytes, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func parseJsonFile(bytes []byte) ([]IntakeError, error) {
	var intakeErrors []IntakeError
	err := json.Unmarshal(bytes, &intakeErrors)

	if err != nil {
		return nil, err
	}

	return intakeErrors, nil
}

func errorByCode(code string, zone string) (*IntakeError, error) {

	bytes, err := readJsonFile("service")

	if err != nil {
		return nil, err
	}

	intakeErrors, err := parseJsonFile(bytes)

	if err != nil {
		return nil, err
	}

	idx := slices.IndexFunc(intakeErrors, func(intakeError IntakeError) bool {
		return intakeError.Code == code
	})

	if idx < 0 {
		return nil, errors.New("IntakeErrors out of range")
	}

	return &intakeErrors[idx], nil
}
