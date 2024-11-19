package structes

type Welcome struct {
	Message string `json:"message"`
	HelperJson
}

func (h Welcome) ToJson() []byte {
	return h.HelperJson.ToJson(h)
}
