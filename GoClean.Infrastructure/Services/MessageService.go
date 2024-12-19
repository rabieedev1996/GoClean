package Services

import (
	"GoClean/GoClean.Domain/Enums"
)

type MessageService struct {
}

func (receiver MessageService) GetMessage(code int) string {
	switch code {
	case Enums.MESSAGE_CODE_SUCCESS:

		return "Operation completed successfully."
	case Enums.MESSAGE_ACTIVE_CODE_IS_EMPETY, Enums.MESSAGE_CODE_EXCEPTION:
		return "An unexpected error occurred."
	case Enums.MESSAGE_CODE_VALIDATION_ERROR:
		return "Validation failed for the input data."
	case Enums.MESSAGE_CODE_INVALID_ACTIVATION:
		return "The activation code is invalid."
	case Enums.MESSAGE_CODE_INVALID_OPERATION:
		return "The requested operation is not allowed."
	case Enums.MESSAGE_CODE_SIMILAR_USERNAME_EXIST:
		return "A user with a similar username already exists."
	case Enums.MESSAGE_CODE_UNAUTHORIZED:
		return "Unauthorized access. Please login first."
	case Enums.MESSAGE_PHONE_NOT_ENTERED:
		return "Phone number is not provided."
	case Enums.MESSAGE_PHONE_IS_INCORRECT:
		return "The provided phone number is incorrect."
	case Enums.MESSAGE_FIRSTNAME_IS_EMPETY:
		return "First name cannot be empty."
	case Enums.MESSAGE_LASTNAME_IS_EMPETY:
		return "Last name cannot be empty."
	default:
		return "Unknown error code."
	}
}
