package dto

type RequestRole struct {
	ID   string `json:"id"`
	Role string `json:"role_id"`
}

type ResponseRole struct {
	ID   string `json:"id"`
	Role string `json:"role_id"`
}

type UpdateResponseRole struct {
	ID   string `json:"id"`
	Role string `json:"role_id"`
}
