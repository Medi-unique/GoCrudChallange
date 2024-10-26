package data

import (
	"errors"

	"github.com/go-playground/validator/v10"

	"example.com/person_crud/models"
)

var Data = map[string]models.Person{
	"1": {
		ID:      "1",
		Name:    "Person 1",
		Age:     20,
		Hobbies: []string{"Reading", "Swimming"},
	},
	"2": {
		ID:      "2",
		Name:    "Person 2",
		Age:     25,
		Hobbies: []string{"Cycling", "Hiking"},
	},
	"3": {
		ID:      "3",
		Name:    "Person 3",
		Age:     30,
		Hobbies: []string{"Cooking", "Gardening"},
	},
	"4": {
		ID:      "4",
		Name:    "Person 4",
		Age:     22,
		Hobbies: []string{"Painting", "Traveling"},
	},
	"5": {
		ID:      "5",
		Name:    "Person 5",
		Age:     28,
		Hobbies: []string{"Photography", "Running"},
	},
	"6": {
		ID:      "6",
		Name:    "Person 6",
		Age:     35,
		Hobbies: []string{"Writing", "Fishing"},
	},
	"7": {
		ID:      "7",
		Name:    "Person 7",
		Age:     27,
		Hobbies: []string{"Dancing", "Skiing"},
	},
	"8": {
		ID:      "8",
		Name:    "Person 8",
		Age:     32,
		Hobbies: []string{"Gaming", "Yoga"},
	},
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}


func AddPerson(person models.Person) (models.Person, error) {
	err := validate.Struct(person)
	if err != nil {
		return models.Person{}, errors.New("validation failed: " + err.Error())
	}
	_, exists := Data[person.ID]
	if !exists {
		Data[person.ID] = person
		return person, nil
	}

	return models.Person{}, errors.New("a person with this ID already exists")
}

func GetAllPersons() map[string]models.Person {
	return Data
}

func GetPerson(id string) (models.Person, error) {
	val, exists := Data[id]
	if !exists {
		return models.Person{}, errors.New("no person with this ID")
	}
	return val, nil
}

func DeletePerson(id string) (models.Person, error) {
	if val, ok := Data[id]; ok {
		delete(Data, id)
		return val, nil
	}
	return models.Person{}, errors.New("no such person ID")
}

func UpdatePerson(person models.Person, id string) (models.Person, error) {
	if val, ok := Data[id]; ok {
		if len(person.Hobbies) > 0 {
			val.Hobbies = person.Hobbies
		}
		if person.Name != "" {
			val.Name = person.Name
		}
		if person.Age > 0 {
			val.Age = person.Age
		}

		Data[id] = val
		return Data[id], nil
	}
	return models.Person{}, errors.New("no person with such ID")
}
