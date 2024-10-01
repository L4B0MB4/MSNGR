package custom_error

type ForwardFailedError struct {
}

func (f *ForwardFailedError) Error() string {
	return "Forwarding to provider(s) failed"
}

func NewForwardFailedError() error {
	return &ForwardFailedError{}
}
