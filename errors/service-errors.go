package errors

func SEC(code string) *IntakeError {

	intakeError, err := errorByCodeOrSlug(code, serviceZone)

	if err != nil {
		return UnknownError()
	}

	return intakeError
}

func SES(slug string) *IntakeError {

	intakeError, err := errorByCodeOrSlug(slug, serviceZone)

	if err != nil {
		return UnknownError()
	}

	return intakeError
}
