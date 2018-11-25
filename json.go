package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Fname string
	Lname string
	Age   int
	// Address	// will be treated as promoted field
	/* Hence output will be
		{
	        "Fname": "Sumeet",
	        "Lname": "Sarkar",
	        "Age": 28,
	        "street": "22nd Street",
	        "city": "Bangalore",
	        "pin": 560001
	}
	*/
	Address Address
	/* Address fields will be in Address block
		{
	        "Fname": "Sumeet",
	        "Lname": "Sarkar",
	        "Age": 28,
	        "Address": {
	                "street": "22nd Street",
	                "city": "Bangalore",
	                "pin": 560001
	        }
	}
	*/
}

type Address struct {
	Street  string `json:"street"` // custom json field naming
	City    string `json:"city"`
	Pincode int    `json:"pin"`
}

func main() {
	person1 := &Person{
		"Sumeet",
		"Sarkar",
		28,
		Address{
			"22nd Street",
			"Bangalore",
			560001,
		},
	}
	// Note: small case fields, unexported fields will not be marshalled
	// Marshall Person struct to JSON string
	b, err := json.MarshalIndent(person1, "", "\t")
	if err != nil {
		panic("Error marshalling struct person")
	}
	fmt.Println(string(b))

	// Unmarshall the JSON to Person struct
	var person2 Person
	json.Unmarshal(b, &person2)
	fmt.Println(person2) // {Sumeet Sarkar 28 {22nd Street Bangalore 560001}}
}
