package model

type Result struct {
	Data string `json:"data"`
}

type Site struct {
	Name string `json:"name"`
	CreatedAt string `json:"created_at"`
}

type Server struct {
	Name string `json:"name"`
	IPAddress string `json:"ipaddress"`
	Status string `json:"statue"`
}
