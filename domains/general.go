package domains

type Error struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type GetApi struct {
	Status string      `json:"status,omitempty"`
	Data   interface{} `json:"data,omitempty"`
	Count  int64       `json:"count,omitempty"`
}
