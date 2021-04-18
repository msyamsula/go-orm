package jsonResponse

type Person struct {
	Id        uint   `json:"id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	AccountId uint   `json:"account_id"`
}
