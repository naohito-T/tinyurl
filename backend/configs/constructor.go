package configs

const (
	// ApplicationPort is the port the application listens on.
	ApplicationPort = ":6500"
	// OpenAPITitle is the title of the OpenAPI spec.
	OpenAPITitle = "TinyURL API"
	// OpenAPIVersion is the version of the OpenAPI spec.
	OpenAPIVersion = "1.0.0"
	// OpenAPIServerPath is the base URL for the OpenAPI spec.
	OpenAPIDocServerPath = "http://localhost:6500/api/v1"
)

// OperationID: このAPI操作の一意の識別子。これは、API内で操作を参照する際に使用されます。
// Method: HTTPメソッドを指定します。この例では http.MethodGet が使われており、これはHTTPのGETリクエストを示します。
// Path: エンドポイントのURLパスを指定します。ここでは "/greeting/{name}" となっており、{name} はパスパラメータを表しています。
// Summary: 短い説明文です。APIのドキュメントに表示され、APIの目的を簡潔に説明します。
// Description: APIエンドポイントの詳細な説明です。ここでは操作の詳細や動作についての追加情報を提供します。
// Tags: このAPI操作に関連付けられたタグのリストです。これにより、APIドキュメント内で類似の操作をグループ化することができます。

// huma.Register(app, huma.Operation{
// 	OperationID: "health",
// 	Method:      http.MethodGet,
// 	Path:        Router.Health,
// 	Summary:     "Health Check",
// 	Description: "Check the health of the service.",
// 	Tags:        []string{"Public"},
// }, func(_ context.Context, _ *HealthCheckParams) (*HealthCheckQuery, error) {
// 	resp := &HealthCheckQuery{
// 		Body: struct{
// 			Message string `json:"message,omitempty" example:"Hello, world!" doc:"Greeting message"`
// 		}{
// 			Message: "ok",
// 		},
// 	}
// 	fmt.Printf("Health Check: %v\n", resp.Body.Message)
// 	return resp, nil
// })
