package dao //data access object

//action yang akan ada di model
type ActionModel interface{
	Create() interface{}
	FindAll() interface{}
	FindOne(id int) interface{}
	Update() interface{}
	Delete() interface{}
}

