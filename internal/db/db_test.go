package db

import (
	"fmt"
	"testing"

	"gorm.io/gorm"
)

func TestInsert(t *testing.T) {
	fmt.Printf("\"lam\": %v\n", "lam")
	user := User{
		Name:  "lam",
		Email: "truonglamthientai321@gmail.com",
		Workout: []Workout{
			{
				Name: "Push up",
			},
			{
				Name: "Pull up",
			},
		},
	}
	Client.Save(&user)

	// read from db

	var users []User
	Client.Preload("Workout").Find(&users)
	// if len(users) != 1 || len(users[0].Workout) != 2 t.Error("failed to insert")

	if len(users) != 1 {
		t.Error("failed to insert")
	}

	if len(users[0].Workout) != 2 {
		t.Error("failed to insert")
	}
	fmt.Printf("users: %+v\n", users)

	// update

	users[0].Workout[0].Name = "Push up 2"

	Client.Session(&gorm.Session{FullSaveAssociations: true}).Save(&users)

	var users2 []User
	Client.Preload("Workout").Find(&users2)

	if users2[0].Workout[0].Name != "Push up 2" {
		t.Error("failed to update")
	}
	fmt.Printf("users2: %+v\n", users2)

	Client.Migrator().DropTable(&User{}, &Workout{}, &Exercise{})
}
