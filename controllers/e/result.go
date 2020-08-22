package e

type Result struct {
	Code int
	Msg  string
	data interface{}
}

func SetResult(code int, Msg string, data interface{}) *Result {
	return &Result{
		Code: code,
		Msg:  Msg,
		data: data,
	}
}
