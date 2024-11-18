package db

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Username string
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
	UserID   uint
	Exercise []Exercise
}

type Exercise struct {
	gorm.Model
	Name       string
	Repetition int
	Set        int
	Weight     int
	WorkoutID  uint
}

func DBInit() *gorm.DB {
	dns := "host=localhost user=gorm password=gorm dbname=gorm port=5432 sslmode=disable TimeZone=Europe/Moscow"
	var er error
	Client, er := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if er != nil {
		panic(er)
	}
	// Client.Migrator().DropTable(&User{}, &Workout{}, &Exercise{})
	Client.AutoMigrate(&User{}, &Workout{}, &Exercise{})
	return Client
}
