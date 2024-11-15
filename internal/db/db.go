package db

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Client *gorm.DB

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Workout  []Workout
}

type Workout struct {
	gorm.Model
	Name     string
	Date     time.Time
	Finish   bool
	Comment  string
	UserID   int
	Exercise []Exercise
}

type Exercise struct {
	gorm.Model
	Name       string
	Repetition int
	Set        int
	Weight     int
	WorkoutID  int
}

func init() {
	dns := "host=localhost user=gorm password=gorm dbname=gorm port=5432 sslmode=disable TimeZone=Europe/Moscow"
	var er error
	Client, er = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if er != nil {
		panic(er)
	}
	// Client.Migrator().DropTable(&User{}, &Workout{}, &Exercise{})
	Client.AutoMigrate(&User{}, &Workout{}, &Exercise{})
}
