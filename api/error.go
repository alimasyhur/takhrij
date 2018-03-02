package api

//ErrorResp struct returns code and message of error
type ErrorResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
