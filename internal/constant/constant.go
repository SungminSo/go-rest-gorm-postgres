package constant

const (
	PostgresConnStr = "sslmode=disable host=%s port=%d users=%s dbname=%s password=%s"
	ConnectingError = "errors connecting postgresql db. connection string: "

	APPLIED  = "applied"
	APPROVED = "approved"
	REJECTED = "rejected"

	NotFoundStr     = "not found"
	BadRequestStr   = "invalid"
	UnauthorizedStr = "unauthorized"
	ConflictedStr   = "already"
)
