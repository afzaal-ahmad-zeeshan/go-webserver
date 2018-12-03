package data

// Person is a mock entity.
type Person struct {
	name string
	age  int
}

// Database is a mock ORM around the Person type.
type Database struct {
	people []Person
}

var database = Database{}
