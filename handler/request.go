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
	NationalID   string `json:"nationalId"`
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
	if r.Name == "" && r.NationalID == "" && r.Phone == "" && r.Email == "" && r.Adress == "" && r.Number <= 0 && r.Neighborhood == "" && r.CEP == "" {
		return fmt.Errorf("O Body da requisição está vazio ou mal formado")
	}

	if r.Name == "" {
		return errParamIsRequired("`Nome`", "string")
	}

	if r.NationalID == "" {
		return errParamIsRequired("NationalID", "string")
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
	NationalID   string `json:"nationalId"`
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
	if r.Name == "" && r.NationalID == "" && r.Phone == "" && r.Email == "" && r.Adress == "" && r.Number <= 0 && r.Neighborhood == "" && r.CEP == "" && r.Password == "" && r.Picture == "" {
		return fmt.Errorf("O Body da requisição está vazio ou mal formado")
	}

	if r.Name == "" {
		return errParamIsRequired("`Nome`", "string")
	}

	if r.NationalID == "" {
		return errParamIsRequired("NationalID", "string")
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
	NationalID   string `json:"nationalId"`
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
	if r.Name != "" || r.NationalID != "" || r.Phone != "" || r.Email != "" || r.Address != "" || r.Number > 0 || r.Neighborhood != "" || r.CEP != "" || r.City != "" || r.State != "" || r.Country != "" || r.Password != "" || r.Picture != "" {
		return nil
	}
	// If none of the fields were provided, return falsy
	return fmt.Errorf("Passe pelo menos um campo na request")
}

// CLIENT REQUESTS
type CreateClientRequest struct {
	Name         string `json:"name"`
	NationalID   string `json:"nationalId"`
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
	if r.Name == "" && r.NationalID == "" && r.Phone == "" && r.Email == "" && r.Adress == "" && r.Number <= 0 && r.Neighborhood == "" && r.CEP == "" {
		logger.ErrorF("The body that broke: %+v", r)
		return fmt.Errorf("O Body da requisição está vazio ou mal formado")
	}

	if r.Name == "" {
		return errParamIsRequired("`Nome`", "string")
	}

	if r.NationalID == "" {
		return errParamIsRequired("NationalID", "string")
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
	NationalID   string `json:"nationalId"`
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
	if r.Name != "" || r.NationalID != "" || r.Phone != "" || r.Email != "" || r.Address != "" || r.Number > 0 || r.Neighborhood != "" || r.CEP != "" || r.City != "" || r.State != "" || r.Country != "" {
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

// Product REQUESTS
type CreateProductRequest struct {
	Title       string  `json:"title"`
	Price       float32 `json:"price"`
	Code        string  `json:"code"`
	Description string  `json:"description"`
	CategoryID  int     `json:"categoryId"`
}

func (r *CreateProductRequest) Validate() error {
	if r.Title == "" {
		return errParamIsRequired("`Titulo`", "string")
	}
	if r.Price <= 0 {
		return errParamIsRequired("`Price`", "float")
	}
	if r.Code == "" {
		return errParamIsRequired("`Codigo`", "string")
	}
	if r.Description == "" {
		return errParamIsRequired("`Descrição`", "string")
	}
	if r.CategoryID <= 0 {
		return errParamIsRequired("`Categoria`", "id")
	}

	return nil
}

type UpdateProductRequest struct {
	Title       string  `json:"title"`
	Price       float32 `json:"price"`
	Code        string  `json:"code"`
	Description string  `json:"description"`
	CategoryID  int     `json:"categoryId"`
}

func (r *UpdateProductRequest) Validate() error {
	// IF ANY EXISTS IS TRUE
	if r.Title != "" || r.Price > 0 || r.Code != "" || r.Description != "" || r.CategoryID > 0 {
		return nil
	}
	// If none of the fields were provided, return falsy
	return fmt.Errorf("Passe pelo menos um campo na request")
}

// Stock REQUESTS
type CreateStockRequest struct {
	Quantity   int     `json:"quantity"`
	PriceOfBuy float32 `json:"priceofbuy"`
	ProductID  int     `json:"productId"`
	SupplierID int     `json:"supplierId"`
}

func (r *CreateStockRequest) Validate() error {
	if r.Quantity <= 0 {
		return errParamIsRequired("`Quantidade`", "int")
	}
	if r.PriceOfBuy <= 0 {
		return errParamIsRequired("`Valor de Compra`", "float")
	}
	if r.ProductID <= 0 {
		return errParamIsRequired("`Produto`", "produto")
	}
	if r.SupplierID <= 0 {
		return errParamIsRequired("`Fornecedor`", "fornecedor")
	}

	return nil
}

type UpdateStockRequest struct {
	Quantity int `json:"quantity"`
}

func (r *UpdateStockRequest) Validate() error {
	// IF ANY EXISTS IS TRUE
	if r.Quantity >= 0 {
		return nil
	}
	// If none of the fields were provided, return falsy
	return fmt.Errorf("Passe pelo menos um campo na request")
}

// SUPPLIER REQUESTS
type CreateSupplierRequest struct {
	Name         string `json:"name"`
	NationalID   string `json:"nationalId"`
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

func (r *CreateSupplierRequest) Validate() error {
	if r.Name == "" && r.NationalID == "" && r.Phone == "" && r.Email == "" && r.Adress == "" && r.Number <= 0 && r.Neighborhood == "" && r.CEP == "" {
		logger.ErrorF("The body that broke: %+v", r)
		return fmt.Errorf("O Body da requisição está vazio ou mal formado")
	}

	if r.Name == "" {
		return errParamIsRequired("`Nome`", "string")
	}

	if r.NationalID == "" {
		return errParamIsRequired("NationalID", "string")
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

type UpdateSupplierRequest struct {
	Name         string `json:"name"`
	NationalID   string `json:"nationalId"`
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

func (r *UpdateSupplierRequest) Validate() error {
	// IF ANY EXISTS IS TRUE
	if r.Name != "" || r.NationalID != "" || r.Phone != "" || r.Email != "" || r.Address != "" || r.Number > 0 || r.Neighborhood != "" || r.CEP != "" || r.City != "" || r.State != "" || r.Country != "" {
		return nil
	}
	// If none of the fields were provided, return falsy
	return fmt.Errorf("Passe pelo menos um campo na request")
}

// CATEGORY REQUESTS
type CreatePaymentMethodRequest struct {
	Title string  `json:"title"`
	Tax   float32 `json:"tax"`
}

func (r *CreatePaymentMethodRequest) Validate() error {
	if r.Title == "" {
		return errParamIsRequired("`Titulo`", "string")
	}
	if r.Tax < 0 {
		return errParamIsRequired("`Tax`", "Float")
	}

	return nil
}

type UpdatePaymentMethodRequest struct {
	Title string  `json:"title"`
	Tax   float32 `json:"tax"`
}

func (r *UpdatePaymentMethodRequest) Validate() error {
	// IF ANY EXISTS IS TRUE
	if r.Title != "" {
		return nil
	}
	if r.Tax >= 0 {
		return nil
	}
	// If none of the fields were provided, return falsy
	return fmt.Errorf("Passe pelo menos um campo na request")
}

// SELL REQUESTS
type CreateSellRequest struct {
	TotalValue      float32 `json:"totalValue"`
	PaymentMethodID int     `json:"paymentMethodId"`
	Status          string  `json:"status"`
	ClientID        int     `json:"clientId"`
	EmployeeID      int     `json:"employeeId"`

	Products []ProductSell
}

type ProductSell struct {
	Quantity  int `json:"quantity"`
	ProductID int `json:"productId"`
}

func (r *CreateSellRequest) Validate() error {
	if r.TotalValue <= 0 {
		return errParamIsRequired("`Valor Total`", "float")
	}
	if r.PaymentMethodID <= 0 {
		return errParamIsRequired("`Metodo de Pagamento`", "Id do Metodo")
	}
	if r.EmployeeID <= 0 {
		return errParamIsRequired("`Funcionario`", "Id do Funcionário")
	}
	if len(r.Products) <= 0 {
		return errParamIsRequired("`Produto`", "Id do Produto e Quantidade")

	}

	return nil
}

type UpdateSellRequest struct {
	Status string `json:"status"`
}

func (r *UpdateSellRequest) Validate() error {
	// IF ANY EXISTS IS TRUE
	if r.Status != "" {
		return nil
	}

	// If none of the fields were provided, return falsy
	return fmt.Errorf("Passe pelo menos um campo na request")
}

// BILL REQUESTS
type CreateBillRequest struct {
	Title       string  `json:"title"`
	Value       float32 `json:"value"`
	Description string  `json:"description"`
	BillingDay  int     `json:"billingday"`
	IsPaid      *bool   `json:"ispaid"`
}

func (r *CreateBillRequest) Validate() error {
	if r.Title == "" {
		return errParamIsRequired("`Titulo`", "string")
	}
	if r.Value <= 0 {
		return errParamIsRequired("`Valor`", "float")
	}
	if r.Description == "" {
		return errParamIsRequired("`Valor Total`", "float")
	}
	if r.BillingDay <= 0 {
		return errParamIsRequired("`Valor Total`", "float")
	}
	if r.IsPaid == nil {
		return errParamIsRequired("Remote", "boolean")
	}

	return nil
}

type UpdateBillRequest struct {
	Title       string  `json:"title"`
	Value       float32 `json:"value"`
	Description string  `json:"description"`
	BillingDay  int     `json:"billingday"`
	IsPaid      *bool   `json:"ispaid"`
}

func (r *UpdateBillRequest) Validate() error {
	// IF ANY EXISTS IS TRUE
	if r.Title != "" || r.Value <= 0 || r.Description == "" || r.BillingDay <= 0 || r.IsPaid != nil {
		return nil
	}

	// If none of the fields were provided, return falsy
	return fmt.Errorf("Passe pelo menos um campo na request")
}

type PayBillRequest struct {
	IsPaid *bool `json:"ispaid"`
}

func (r *PayBillRequest) Validate() error {
	// IF ANY EXISTS IS TRUE
	if r.IsPaid != nil {
		return errParamIsRequired("Remote", "boolean")
	}

	// If none of the fields were provided, return falsy
	return nil
}
