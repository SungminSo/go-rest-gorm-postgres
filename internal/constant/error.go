package constant

const (
	NotFoundError = "%s not found: %v"

	FindingError  = "errors finding %s: %v"
	SavingError   = "errors saving %s: %v"

	MissingAccessTokenError = "missing accessToken at header"
	InvalidAccessTokenError = "invalid accessToken at header"
	CorruptedJWTError       = "jwt token corrupted"
	InValidJWTError			= "jwt is not valid"
)
