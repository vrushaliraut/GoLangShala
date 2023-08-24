package main

import "time"

type Status string

// Define Enum
// Define model
// Use enums in goorm operations

const (
	Pending  Status = "pending"
	Approved Status = "approved"
	Rejected Status = "rejected"
)

// model
type User struct {
	ID        uint `gorm:"primary_key"`
	Name      string
	Status    Status
	CreatedAt time.Time
	UpdatedAt time.Time
}

// use enum in gorm operations
func enums_with_goorm() {
	//create a user with status pending
	/*
		user := User{Name: "Vrushali", Status: Pending}
		db.Create(&user)

		//query all users with status approve
		var users []User
		db.Where("status = ?", Approved).Find(&users)

		// Update a user's status to "rejected"
		db.Model(&user).Update("status", Rejected)

		// Delete all users with status "pending"
		db.Where("status = ?", Pending).Delete(User{})


	*/
}
