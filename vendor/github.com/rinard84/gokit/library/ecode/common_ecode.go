package ecode

// All common ecode
var (
	OK = add(0)

	NotModified        = add(-304)
	TemporaryRedirect  = add(-307)
	RequestErr         = add(-400)
	Unauthorized       = add(-401)
	AccessDenied       = add(-403)
	NothingFound       = add(-404)
	MethodNotAllowed   = add(-405)
	Conflict           = add(-409)
	Canceled           = add(-498)
	ServerErr          = add(-500)
	ServiceUnavailable = add(-503)
	Deadline           = add(-504)
	LimitExceed        = add(-509)

	JwtCodeSecretUnmatched = Error(401, "Token invalid") // "secret unmatched"
	JwtCodeTokenExpired    = Error(402, "Token expire")  // "token expired"
	JwtCodeNotAllowedAud   = Error(403, "Token invalid") // "not allowed aud"
	JwtCodeNotAllowedIss   = Error(404, "Token invalid") // "not allowed iss"
	JwtCodeParseClaim      = Error(405, "Token invalid") // "couldn't parse claims"

	InvalidInput = Error(406, "Invalid input")
)
