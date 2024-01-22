package bank

func Ping() string {
	return "TEST RESPONSE"
}

type Customer struct {
	Name    string
	Address string
	Phone   string
}

type Account struct {
	Customer
	Number  int32
	Balance float64
}
