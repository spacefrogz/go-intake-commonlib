package errors

func PE(code string) *IntakeError {

	intakeError, err := errorByCode(code, presenterZone)

	if err != nil {
		return UnknownError()
	}

	return intakeError
}
