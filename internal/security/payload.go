package security

type SessionPayload struct {
	IP      string `json:"ip"`
	Port    string `json:"port"`
	Session string `json:"session"`
	Expiry  int64  `json:"exp"`
	Pwd     string `json:"pwd"`
}
