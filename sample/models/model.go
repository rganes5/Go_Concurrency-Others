package model

type Post struct {
	Body   string `json:"body"`
	UserID int    `json:"userId"`
}

type Req struct {
	Amount float64 `json:"amount"`
}

type Res struct {
	Discount float64 `json:"discount"`
	Suc      string  `json:"suc"`
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
