package fieldvalidation

type FieldValidationRule string

const (
	IsRequired   FieldValidationRule = "is_required"
	IsNumeric    FieldValidationRule = "is_numeric"
	IsEmail      FieldValidationRule = "is_email"
	IsChecked    FieldValidationRule = "is_checked"
	IsNotChecked FieldValidationRule = "is_not_checked"
)
