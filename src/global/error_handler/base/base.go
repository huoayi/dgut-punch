package error_handler

type GinError struct {
	Status  int    `json: "-"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (err GinError) Error() string {
	return err.Message
}
