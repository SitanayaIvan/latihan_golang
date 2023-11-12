package domains

type Success struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Count  int64       `json:"count,omitempty"`
}

type Error struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
