package param

type LoginParam struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Id       string `json:"id"`
	Value    string `json:"code"`
}
