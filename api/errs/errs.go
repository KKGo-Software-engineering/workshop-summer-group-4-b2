package errs

type ErrResponse struct {
	Message string `json:"message"`
}

func Build(err error) ErrResponse {
	return ErrResponse{
		Message: err.Error(),
	}
}
