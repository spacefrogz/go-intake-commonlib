package errors

func PE(code string) *IntakeError {

	intakeError, err := errorByCode(code, "presenter")

	if err != nil {
		return UnknownError()
	}

	return intakeError
}
