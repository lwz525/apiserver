package errno

var (
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{10001, "Internal server error."}
	ErrBind             = &Errno{10002, "Error occurred while binding the request body to the struct."}
	ErrUserNotFound     = &Errno{20102, "The user was not found."}
	ErrValidation       = &Errno{Code: 20001, Message: "Validation failed."}
	ErrDatabase         = &Errno{Code: 20002, Message: "Database error."}
	ErrToken            = &Errno{Code: 20003, Message: "Error occurred while signing the JSON web token."}

	// user errors
	ErrEncrypt           = &Errno{Code: 20101, Message: "Error occurred while encrypting the user password."}
	ErrTokenInvalid      = &Errno{Code: 20103, Message: "The token was invalid."}
	ErrPasswordIncorrect = &Errno{Code: 20104, Message: "The password was incorrect."}

)
