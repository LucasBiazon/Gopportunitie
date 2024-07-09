package handler

import "fmt"

func ErrorParameIsRequied(param string, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", param, typ)
}

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == " " && r.Company == " " && r.Location == " " && r.Remote == nil && r.Link == " " && r.Salary <= 0 {
		return ErrorParameIsRequied("request", "object")
	}
	if r.Role == " " {
		return ErrorParameIsRequied("role", "string")
	}
	if r.Company == " " {
		return ErrorParameIsRequied("company", "string")
	}
	if r.Location == " " {
		return ErrorParameIsRequied("location", "string")
	}

	if r.Remote == nil {
		return ErrorParameIsRequied("remote", "bool")
	}

	if r.Link == " " {
		return ErrorParameIsRequied("link", "string")
	}

	if r.Salary <= 0 {
		return ErrorParameIsRequied("salary", "int")

	}
	return nil
}
