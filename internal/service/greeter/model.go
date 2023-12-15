package greeter

type ReqGreet struct {
	Message string `json:"message"`
}

type ResGreet struct {
	Message string `json:"message"`
}
