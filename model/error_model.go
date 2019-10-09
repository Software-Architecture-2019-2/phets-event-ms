package model

// Error model to return
type Error struct {
	Code        int              `json:"code"`
	Message     string           `json:"message"`
	Description ErrorDescription `json:"description"`
}

// ErrorDescription with the details of the error
type ErrorDescription struct {
	Error string `json:"error"`
}
