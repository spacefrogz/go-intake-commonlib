package errors

func SE(code string) *IntakeError {

	intakeError, err := errorByCode(code, serviceZone)

	if err != nil {
		return UnknownError()
	}

	return intakeError
}
