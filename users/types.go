package users

type User struct {
	Id      int64  `json:"id"`
	Name    string `json:"name,omitempty"`
	SurName string `json:"surname,omitempty"`
	Age     byte   `json:"age,omitempty"`
}

func (User) SwaggerDoc() map[string]string {
	return map[string]string{
		"":        "User doc",
		"id":      "Unique ID",
		"name":    "Name of the user",
		"surname": "Surname of the user",
		"age":     "Age of the user",
	}
}
