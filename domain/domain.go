package domain

type Location struct {
	IsHQ        bool     `json:"is_hq"`
	Country     string   `json:"country"`
	City        string   `json:"city"`
	PostalCode  string   `json:"postal_code"`
	Street      []string `json:"street"`
	Description string   `json:"description"`
	Area        string   `json:"area"`
}

type EmployeeCountRange struct {
	From int `json:"from"`
	To   int `json:"to"`
}

type Company struct {
	LinkedInID              string             `json:"linkedin_id"`
	URN                     string             `json:"urn"`
	Name                    string             `json:"name"`
	PublicID                string             `json:"public_id"`
	Description             string             `json:"description"`
	LinkedInURL             string             `json:"linkedin_url"`
	Hashtags                []string           `json:"hashtags"`
	FollowCount             int                `json:"follow_count"`
	Locations               []Location         `json:"locations"`
	Tagline                 string             `json:"tagline"`
	WebsiteURL              string             `json:"website_url"`
	Phone                   string             `json:"phone"`
	FoundedDate             string             `json:"founded_date"`
	EstimatedEmployeeCount  int                `json:"estimated_employee_count"`
	EmployeeCountRange      EmployeeCountRange `json:"employee_count_range"`
	Industry                []string           `json:"industry"`
}

type Tenure struct {
	Years  int `json:"years"`
	Months int `json:"months"`
}

type CompanyData struct {
	Name        string `json:"name"`
	LinkedInID  string `json:"linkedin_id"`
	Description string `json:"description"`
	Industry    string `json:"industry"`
}

type CurrentPosition struct {
	CompanyData        CompanyData `json:"company_data"`
	Role               string      `json:"role"`
	Location           string      `json:"location"`
	TenureAtRole       Tenure      `json:"tenure_at_role"`
	TenureAtCompany    Tenure      `json:"tenure_at_company"`
}

type Employee struct {
	LinkedInID      string          `json:"linkedin_id"`
	PublicID        string          `json:"public_id"`
	URN             string          `json:"urn"`
	LinkedInURL     string          `json:"linkedin_url"`
	Name            string          `json:"name"`
	FirstName       string          `json:"first_name"`
	LastName        string          `json:"last_name"`
	Location        string          `json:"location"`
	Headline        string          `json:"headline"`
	Summary         string          `json:"summary"`
	Premium         bool            `json:"premium"`
	CurrentPosition CurrentPosition `json:"current_position"`
}

type EnrichmentRequest struct {
	CompanyLinkedInURL string `json:"company_linkedin_url,omitempty"`
	CompanyLinkedInID  string `json:"company_linkedin_id,omitempty"`
}

type EnrichmentResponse struct {
	Company   Company    `json:"company"`
	Employees []Employee `json:"employees"`
}

type APIResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Error  string      `json:"error,omitempty"`
}

type CompanyAPIResponse struct {
	Status string `json:"status"`
	Data   struct {
		Company Company `json:"company"`
	} `json:"data"`
	Error string `json:"error,omitempty"`
}

type EmployeeAPIResponse struct {
	Status string `json:"status"`
	Data   struct {
		CompanyLinkedInID string     `json:"company_linkedin_id"`
		Cursor            *string    `json:"cursor,omitempty"`
		Employees         []Employee `json:"employees"`
	} `json:"data"`
	Error string `json:"error,omitempty"`
}