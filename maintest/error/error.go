package error

type Error struct {
	code int    `json:"code"`
	msg  string `json:"msg"`
}
