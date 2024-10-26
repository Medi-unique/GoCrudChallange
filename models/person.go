package models

type Person struct {
    ID      string   `json:"id" validate:"required"`
    Name    string   `json:"name" validate:"required,min=1"`
    Age     int      `json:"age" validate:"required,gt=0"`
    Hobbies []string `json:"hobbies" validate:"required"`
}
