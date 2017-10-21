package main

type Employee struct {
	Id   int    `xorm:"int autoincr not null pk 'employee_id'"`
	Name string `xorm:"varchar(255) not null 'employee_name'"`
}

func (e Employee) TableName() string {
	return "employee"
}
