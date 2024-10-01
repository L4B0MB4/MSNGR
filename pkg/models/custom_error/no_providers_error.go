package custom_error

type NoProvidersError struct {
}

func (f *NoProvidersError) Error() string {
	return "No providers found to forward to"
}

func NewNoProvidersError() error {
	return &NoProvidersError{}
}
