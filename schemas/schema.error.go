package schemas

type SchemaDatabaseError struct {
	Type string
	Code int
}
type SchemaErrorResponse struct {
	StatusCode int         `json:"statusCode"`
	Method     string      `json:"method"`
	Error      interface{} `json:"error"`
}
