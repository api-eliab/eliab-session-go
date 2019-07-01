package main

// LoginResponse ...
type ResponseLogin struct {
	User ResponseUser `json:"user"`
}

// ResponseUser ...
type ResponseUser struct {
	ID        int64         `json:"id"`
	FirstName string        `json:"firstName"`
	LastName  string        `json:"lastName"`
	Email     string        `json:"email"`
	Phone     string        `json:"phone"`
	Address   string        `json:"address"`
	Sons      []ResponseSon `json:"sons"`
}

// ResponseSon ...
type ResponseSon struct {
	ID        int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Avatar    int64  `json:"avatar"`
}
