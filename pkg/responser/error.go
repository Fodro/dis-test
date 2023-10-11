package pkg

type Error struct {
	Code  int         `json:"code"`
	Error interface{} `json:"error"`
}
