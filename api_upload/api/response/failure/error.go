package failure

type ServerError struct {
	Reason string `json:"reason"`
}

func (s ServerError) Error() string {
	return s.Reason
}
