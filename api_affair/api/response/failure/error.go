package failure

type ErrorResponse struct {
	Reason string `json:"reason"`
}

func (s ErrorResponse) Error() string {
	return s.Reason
}
