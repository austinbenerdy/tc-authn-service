package errors

type AuthFailedError struct{}

func (err *AuthFailedError) Error() string {
	return "Auth Failed"
}
