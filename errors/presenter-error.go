package errors

func PEC(code string) *IntakeError {

	intakeError, err := errorByCodeOrSlug(code, presenterZone)

	if err != nil {
		return UnknownError()
	}

	return intakeError
}

func PES(slug string) *IntakeError {

	intakeError, err := errorByCodeOrSlug(slug, presenterZone)

	if err != nil {
		return UnknownError()
	}

	return intakeError
}
