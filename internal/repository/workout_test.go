package repository

import (
	"fmt"
	"testing"
	"time"

	"github.com/kaitokid2302/Workout-Tracker/internal/db"
	"github.com/stretchr/testify/assert"
)

func SetupTest() {
	client := db.DBInit()
	client.Migrator().DropTable(&db.User{}, &db.Workout{}, &db.Exercise{})
	client.AutoMigrate(&db.User{}, &db.Workout{}, &db.Exercise{})

	// user1: name = lam, workout [workout1, workout2], workout1: exercise [exercise1, exercise2], workout2: exercise [exercise3]
	// user2: name = linh, workout [workout3, workout4], workout3: exercise [exercise4, exercise5], workout4: exercise [exercise6]
	user1 := db.User{
		Name:     "Lam",
		Username: "lamdeptrai",
		Email:    "truonglamthientai321@gmail.com",
		Password: "123456",
		Workout: []db.Workout{
			{
				Name:    "Morning Workout",
				Date:    time.Date(2025, time.January, 15, 8, 0, 0, 0, time.FixedZone("MSK", 3*60*60)), // 8:00 MSK
				Finish:  false,
				Comment: "Need to increase weight next time",
				Exercise: []db.Exercise{
					{
						Name:       "Bench Press",
						Repetition: 10,
						Set:        3,
						Weight:     60,
					},
					{
						Name:       "Shoulder Press",
						Repetition: 12,
						Set:        3,
						Weight:     40,
					},
				},
			},
			{
				Name:    "Evening Workout",
				Date:    time.Date(2025, time.January, 15, 18, 30, 0, 0, time.FixedZone("MSK", 3*60*60)), // 18:30 MSK
				Finish:  true,
				Comment: "Great session",
				Exercise: []db.Exercise{
					{
						Name:       "Squat",
						Repetition: 8,
						Set:        4,
						Weight:     100,
					},
				},
			},
		},
	}

	user2 := db.User{
		Name:     "Linh",
		Username: "linhxinh",
		Email:    "linh@gmail.com",
		Password: "654321",
		Workout: []db.Workout{
			{
				Name:    "Leg Day",
				Date:    time.Date(2025, time.February, 1, 10, 15, 0, 0, time.FixedZone("MSK", 3*60*60)), // 10:15 MSK
				Finish:  true,
				Comment: "Legs are shaking",
				Exercise: []db.Exercise{
					{
						Name:       "Deadlift",
						Repetition: 5,
						Set:        5,
						Weight:     80,
					},
					{
						Name:       "Leg Press",
						Repetition: 12,
						Set:        3,
						Weight:     120,
					},
				},
			},
			{
				Name:    "Upper Body",
				Date:    time.Date(2025, time.February, 1, 17, 45, 0, 0, time.FixedZone("MSK", 3*60*60)), // 17:45 MSK
				Finish:  false,
				Comment: "Need to work on form",
				Exercise: []db.Exercise{
					{
						Name:       "Pull Ups",
						Repetition: 8,
						Set:        4,
						Weight:     0,
					},
				},
			},
		},
	}

	client.Save(&user1)
	client.Save(&user2)
}

func TestRepository(t *testing.T) {
	SetupTest()
	client := db.DBInit()

	t.Run("FindUserByUsername", func(t *testing.T) {
		repo := NewWorkoutRepository(client)
		user, err := repo.FindUserByUsername("lamdeptrai")
		// assert nil
		assert.Nil(t, err)
		// asert Name = "Lam"
		assert.Equal(t, "Lam", user.Name)

		fmt.Printf("user: %v\n", user)
	})
}
