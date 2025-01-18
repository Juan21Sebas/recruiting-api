package schema

type Response struct {
	Data   any    `json:"data,omitempty" mask:"struct"`
	Result Result `json:"result"`
}

type Result struct {
	Details []Detail `json:"details"`
	Source  string   `json:"source"`
}

type Detail struct {
	InternalCode string `json:"internalCode"`
	Message      string `json:"message"`
	Detail       string `json:"detail"`
}

type CandidatesGetResponse struct {
	Name           string  `json:"name"`
	Email          string  `json:"email"`
	Gender         string  `json:"gender"`
	SalaryExpected float64 `json:"salary_expected"`
}

type CandidatesUpdateResponse struct {
	Name           string  `json:"name"`
	Email          string  `json:"email"`
	Gender         string  `json:"gender"`
	SalaryExpected float64 `json:"salary_expected"`
}
