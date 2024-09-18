package query

type HealthCheckQuery struct {
	CheckDB bool `query:"q" doc:"Optional DynamoDB check parameter"`
}
