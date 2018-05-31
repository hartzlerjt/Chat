package chat

// Struct for a chat message
type Message struct {
	From int    `json:"from"`
	To   int    `json:"to"`
	Mesg []byte `json:"mesg"`
}

// Struct that defines the format of a login
type Login struct {
	User     string `json"user"`
	Password string `json:"password"`
}
