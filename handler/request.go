package handler

//CreateOpening
type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() bool {
	if r.Role == "" {
		return false
	}

	if r.Company == "" {
		return false
	}

	if r.Location == "" {
		return false
	}

	if r.Link == "" {
		return false
	}

	if r.Salary == 0 {
		return false
	}

	if r.Remote == nil {
		return false
	}

	if r.Link == "" {
		return false
	}
	return true
}
