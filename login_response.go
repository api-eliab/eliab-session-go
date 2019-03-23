package main

// LoginResponse ...
type LoginResponse struct {
	User struct {
		ID        int    `json:"id"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
		Address   string `json:"address"`
		Sons      []struct {
			ID        int    `json:"id"`
			FirstName string `json:"firstName"`
			LastName  string `json:"lastName"`
			Avatar    int    `json:"avatar"`
		} `json:"sons"`
	} `json:"user"`
}
