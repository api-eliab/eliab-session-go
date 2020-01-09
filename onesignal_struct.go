package main

// MessageRequest doc ...
type MessageRequest struct {
	Message struct {
		Title   string  `json:"title"`
		Message string  `json:"message"`
		Icon    string  `json:"icon"`
		Users   []int64 `json:"users,omitempty"`
	} `json:"message"`
}
