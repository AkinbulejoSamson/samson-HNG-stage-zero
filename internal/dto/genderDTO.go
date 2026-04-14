package dto

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ProcessedData struct {
	Name        string  `json:"name"`
	Gender      string  `json:"gender"`
	Probability float64 `json:"probability"`
	SampleSize  int     `json:"sample_size"`
	IsConfident bool    `json:"is_confident"`
	ProcessedAt string  `json:"processed_at"`
}

type SuccessResponse struct {
	Status string        `json:"status"`
	Data   ProcessedData `json:"data"`
}

type GenderizeRawData struct {
	Name        string  `json:"name"`
	Gender      string  `json:"gender"`
	Probability float64 `json:"probability"`
	Count       int     `json:"count"`
}
