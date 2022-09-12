package errors

func SE(code string) *IntakeError {

	intakeError, err := errorByCode(code, "service")

	if err != nil {
		return UnknownError()
	}

	return intakeError
}
