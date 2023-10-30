package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("O parâmetro: %s (do tipo: %s) é obrigatório", name, typ)
}

// CreateOpening

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
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
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *UpdateOpeningRequest) Validate() error {
	// IF ANY ECISTS IS TRUE
	if r.Role != "" || r.Company != "" || r.Location != "" || r.Remote != nil || r.Link != "" || r.Salary > 0 {
		return nil
	}
	// If none of the fields were provided, return falsy
	return fmt.Errorf("Passe pelo menos um campo na request")
}

// USER REQUESTS
type CreateUserRequest struct {
	Name         string `json:"name"`
	CPF          string `json:"cpf"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	Adress       string `json:"adress"`
	Number       uint   `json:"number"`
	Neighborhood string `json:"neighborhood"`
	CEP          string `json:"cep"`
	City         string `json:"city"`
	State        string `json:"state"`
	Country      string `json:"country"`
}

func (r *CreateUserRequest) Validate() error {
	if r.Name == "" && r.CPF == "" && r.Phone == "" && r.Email == "" && r.Adress == "" && r.Number <= 0 && r.Neighborhood == "" && r.CEP == "" {
		return fmt.Errorf("O Body da requisição está vazio ou mal formado")
	}

	if r.Name == "" {
		return errParamIsRequired("`Nome`", "string")
	}

	if r.CPF == "" {
		return errParamIsRequired("CPF", "string")
	}

	if r.Phone == "" {
		return errParamIsRequired("Telefone", "string")
	}

	if r.Email == "" {
		return errParamIsRequired("Email", "string")
	}

	if r.Adress == "" {
		return errParamIsRequired("Endereço", "string")
	}

	if r.Number <= 0 {
		return errParamIsRequired("Numero", "number")
	}

	if r.Neighborhood == "" {
		return errParamIsRequired("Bairro", "string")
	}

	if r.CEP == "" {
		return errParamIsRequired("Cep", "string")
	}

	return nil
}

// EMPLOYEE REQUESTS
type CreateUserWorkerRequest struct {
	Name         string `json:"name"`
	CPF          string `json:"cpf"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	Adress       string `json:"adress"`
	Number       uint   `json:"number"`
	Neighborhood string `json:"neighborhood"`
	CEP          string `json:"cep"`
	City         string `json:"city"`
	State        string `json:"state"`
	Country      string `json:"country"`
	Password     string `json:"password"`
	Picture      string `json:"picture"`
}

func (r *CreateUserWorkerRequest) Validate() error {
	if r.Name == "" && r.CPF == "" && r.Phone == "" && r.Email == "" && r.Adress == "" && r.Number <= 0 && r.Neighborhood == "" && r.CEP == "" && r.Password == "" && r.Picture == "" {
		return fmt.Errorf("O Body da requisição está vazio ou mal formado")
	}

	if r.Name == "" {
		return errParamIsRequired("`Nome`", "string")
	}

	if r.CPF == "" {
		return errParamIsRequired("CPF", "string")
	}

	if r.Phone == "" {
		return errParamIsRequired("Telefone", "string")
	}

	if r.Email == "" {
		return errParamIsRequired("Email", "string")
	}

	if r.Adress == "" {
		return errParamIsRequired("Endereço", "string")
	}

	if r.Number <= 0 {
		return errParamIsRequired("Numero", "number")
	}

	if r.Neighborhood == "" {
		return errParamIsRequired("Bairro", "string")
	}

	if r.CEP == "" {
		return errParamIsRequired("Cep", "string")
	}
	if r.Password == "" {
		return errParamIsRequired("Cep", "string")
	}
	if r.Picture == "" {
		return errParamIsRequired("Cep", "string")
	}

	return nil
}

type UpdateEmployeeRequest struct {
	Name         string `json:"name"`
	CPF          string `json:"cpf"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	Address      string `json:"adress"`
	Number       uint   `json:"number"`
	Neighborhood string `json:"neighborhood"`
	CEP          string `json:"cep"`
	City         string `json:"city"`
	State        string `json:"state"`
	Country      string `json:"country"`
	Password     string `json:"password"`
	Picture      string `json:"picture"`
}

func (r *UpdateEmployeeRequest) Validate() error {
	// IF ANY EXISTS IS TRUE
	if r.Name != "" || r.CPF != "" || r.Phone != "" || r.Email != "" || r.Address != "" || r.Number > 0 || r.Neighborhood != "" || r.CEP != "" || r.City != "" || r.State != "" || r.Country != "" || r.Password != "" || r.Picture != "" {
		return nil
	}
	// If none of the fields were provided, return falsy
	return fmt.Errorf("Passe pelo menos um campo na request")
}

// CLIENT REQUESTS
type CreateClientRequest struct {
	Name         string `json:"name"`
	CPF          string `json:"cpf"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	Adress       string `json:"adress"`
	Number       uint   `json:"number"`
	Neighborhood string `json:"neighborhood"`
	CEP          string `json:"cep"`
	City         string `json:"city"`
	State        string `json:"state"`
	Country      string `json:"country"`
}

func (r *CreateClientRequest) Validate() error {
	if r.Name == "" && r.CPF == "" && r.Phone == "" && r.Email == "" && r.Adress == "" && r.Number <= 0 && r.Neighborhood == "" && r.CEP == "" {
		logger.ErrorF("The body that broke: %+v", r)
		return fmt.Errorf("O Body da requisição está vazio ou mal formado")
	}

	if r.Name == "" {
		return errParamIsRequired("`Nome`", "string")
	}

	if r.CPF == "" {
		return errParamIsRequired("CPF", "string")
	}

	if r.Phone == "" {
		return errParamIsRequired("Telefone", "string")
	}

	if r.Email == "" {
		return errParamIsRequired("Email", "string")
	}

	if r.Adress == "" {
		return errParamIsRequired("Endereço", "string")
	}

	if r.Number <= 0 {
		return errParamIsRequired("Numero", "number")
	}

	if r.Neighborhood == "" {
		return errParamIsRequired("Bairro", "string")
	}

	if r.CEP == "" {
		return errParamIsRequired("Cep", "string")
	}

	return nil
}

type UpdateClientRequest struct {
	Name         string `json:"name"`
	CPF          string `json:"cpf"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	Address      string `json:"adress"`
	Number       uint   `json:"number"`
	Neighborhood string `json:"neighborhood"`
	CEP          string `json:"cep"`
	City         string `json:"city"`
	State        string `json:"state"`
	Country      string `json:"country"`
}

func (r *UpdateClientRequest) Validate() error {
	// IF ANY EXISTS IS TRUE
	if r.Name != "" || r.CPF != "" || r.Phone != "" || r.Email != "" || r.Address != "" || r.Number > 0 || r.Neighborhood != "" || r.CEP != "" || r.City != "" || r.State != "" || r.Country != "" {
		return nil
	}
	// If none of the fields were provided, return falsy
	return fmt.Errorf("Passe pelo menos um campo na request")
}

// CATEGORY REQUESTS
type CreateCategoryRequest struct {
	Title string `json:"title"`
}

func (r *CreateCategoryRequest) Validate() error {
	if r.Title == "" {
		return errParamIsRequired("`Titulo`", "string")
	}

	return nil
}

type UpdateCategoryRequest struct {
	Title string `json:"title"`
}

func (r *UpdateCategoryRequest) Validate() error {
	// IF ANY EXISTS IS TRUE
	if r.Title != "" {
		return nil
	}
	// If none of the fields were provided, return falsy
	return fmt.Errorf("Passe pelo menos um campo na request")
}
