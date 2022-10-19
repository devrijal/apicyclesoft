package auth

type User struct {
	APIKey    string `json:"api_key"`
	APISecret string `json:"api_secret"`
}

var users = []User{
	{
		APIKey:    "YOUR_API_KEY",
		APISecret: "YOUR_API_SECRET",
	},
}
