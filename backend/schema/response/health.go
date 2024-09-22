package response

type HealthCheckResponse struct {
	Body struct {
		Message string `json:"message"`
	}
}
