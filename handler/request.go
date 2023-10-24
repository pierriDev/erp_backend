package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("O parâmetro: %s (do tipo: %s) é obrigatório", name, typ)
}

// CreateOpening

type CreateOpeningRequest struct {
	Role     string `json: "role"`
	Company  string `json: "company"`
	Location string `json: "location"`
	Remote   *bool  `json: "remote"`
	Link     string `json: "link"`
	Salary   int64  `json: "salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" && r.Company == "" && r.Location == "" && r.Remote == nil && r.Salary <= 0 {
		return fmt.Errorf("O Body da requisição está vazio ou mal formado")
	}

	if r.Role == "" {
		return errParamIsRequired("Role", "string")
	}
	if r.Company == "" {
		return errParamIsRequired("Company", "string")
	}
	if r.Location == "" {
		return errParamIsRequired("Location", "string")
	}
	if r.Link == "" {
		return errParamIsRequired("Link", "string")
	}

	if r.Remote == nil {
		return errParamIsRequired("Remote", "boolean")
	}
	if r.Salary <= 0 {
		return errParamIsRequired("Salary", "int")
	}
	return nil
}

type UpdateOpeningRequest struct {
	Role     string `json: "role"`
	Company  string `json: "company"`
	Location string `json: "location"`
	Remote   *bool  `json: "remote"`
	Link     string `json: "link"`
	Salary   int64  `json: "salary"`
}

func (r *UpdateOpeningRequest) Validate() error {
	// IF ANY ECISTS IS TRUE
	if r.Role != "" || r.Company != "" || r.Location != "" || r.Remote != nil || r.Link != "" || r.Salary > 0 {
		return nil
	}
	// If none of the fields were provided, return falsy
	return fmt.Errorf("Passe pelo menos um campo na request")
}
