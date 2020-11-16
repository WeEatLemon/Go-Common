package EmailModules

type Notice struct {
	Type string      `json:"type"`
	Msg  interface{} `json:"msg"`
}
