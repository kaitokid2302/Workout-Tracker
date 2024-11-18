package repository

import (
	"fmt"
	"testing"
	"time"

	"github.com/kaitokid2302/Workout-Tracker/internal/db"
	"github.com/kaitokid2302/Workout-Tracker/internal/utils"
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
	repo := NewWorkoutRepository(client)

	// FindUserByUsername
	t.Run("FindUserByUsername", func(t *testing.T) {

		user, err := repo.FindUserByUsername("lamdeptrai")
		// assert nil
		assert.Nil(t, err)
		// asert Name = "Lam"
		assert.Equal(t, "Lam", user.Name)

		utils.PrintBlue(fmt.Sprintf("user: %+v\n", user))

		user, err = repo.FindUserByUsername("linhxinh")
		// assert nil
		assert.Nil(t, err)
		// asert Name = "Linh"
		assert.Equal(t, "Linh", user.Name)

		utils.PrintBlue(fmt.Sprintf("user: %+v\n", user))
	})

	// CreateEmptyWorkout

	t.Run("CreateEmptyWorkout", func(t *testing.T) {
		user, _ := repo.FindUserByUsername("lamdeptrai")
		err := repo.CreateEmptyWorkout(user, "Afternoon Workout")
		// assert nil
		assert.Nil(t, err)
		var workout db.Workout
		er := client.Where("name = ?", "Afternoon Workout").First(&workout).Error

		// assert nil
		assert.Nil(t, er)
		// assert Name = "Afternoon Workout"
		assert.Equal(t, "Afternoon Workout", workout.Name)

		// assert userID = 1
		assert.Equal(t, uint(1), workout.UserID)

		utils.PrintBlue(fmt.Sprintf("workout: %+v\n", workout))
	})

	// AddExerciseToWorkout
	t.Run("AddExerciseToWorkout", func(t *testing.T) {
		// find workout Afternoon Workout
		var workout db.Workout
		er := client.Where("name = ?", "Afternoon Workout").First(&workout).Error
		// assert nil
		assert.Nil(t, er)

		exercise := db.Exercise{
			Name:       "Bicep Curl",
			Repetition: 12,
			Set:        3,
			Weight:     20,
		}

		err := repo.AddExerciseToWorkout(&workout, &exercise)
		// assert nil
		assert.Nil(t, err)

		// find exercise Bicep Curl
		var e db.Exercise
		er = client.Where("name = ?", "Bicep Curl").First(&e).Error
		// assert nil
		assert.Nil(t, er)

		// assert workoutID = workout.ID
		assert.Equal(t, workout.ID, e.WorkoutID)
		// assert Name = "Bicep Curl"
		assert.Equal(t, "Bicep Curl", e.Name)

		utils.PrintBlue(fmt.Sprintf("exercise: %+v\n", e))
	})

	// DeleteExercise
	t.Run("DeleteExercise", func(t *testing.T) {
		// Delete exercise Bicep Curl
		var e db.Exercise
		er := client.Where("name = ?", "Bicep Curl").First(&e).Error
		// assert nil
		assert.Nil(t, er)
		// assrt Name = "Bicep Curl"
		assert.Equal(t, "Bicep Curl", e.Name)

		// find workout Afternoon Workout
		var workout db.Workout
		er = client.Where("name = ?", "Afternoon Workout").First(&workout).Error
		// assert nil
		assert.Nil(t, er)

		err := repo.DeleteExercise(&workout, e.ID)
		// assert nil
		assert.Nil(t, err)

		// find exercise Bicep Curl
		e = db.Exercise{}
		er = client.Where("name = ?", "Bicep Curl").First(&e).Error
		// assert not nil
		assert.NotNil(t, er)

		// assert Name = ""
		assert.Equal(t, "", e.Name)
	})

	// EditComment
	t.Run("EditComment", func(t *testing.T) {
		// Morning Workout -> "hello, new comment"
		var workout db.Workout
		er := client.Where("name = ?", "Morning Workout").First(&workout).Error
		// assert nil

		assert.Nil(t, er)
		repo.EditComment(&workout, "hello, new comment")
		workout = db.Workout{}

		er = client.Where("name = ?", "Morning Workout").First(&workout).Error
		// assert nil
		assert.Nil(t, er)

		// assert Comment = "hello, new comment"
		assert.Equal(t, "hello, new comment", workout.Comment)
	})

	// DeleteWorkout
	t.Run("DeleteWorkout", func(t *testing.T) {
		// delete morning workout
		var workout db.Workout
		er := client.Where("name = ?", "Morning Workout").First(&workout).Error
		// assert nil
		assert.Nil(t, er)
		// assert name = "Morning Workout"
		assert.Equal(t, "Morning Workout", workout.Name)

		err := repo.DeleteWorkout(&workout)
		// assert nil
		assert.Nil(t, err)
		// find all workout

		workouts := []db.Workout{}
		client.Find(&workouts)

		// count = 4
		assert.Equal(t, 4, len(workouts))
		t.Logf("workouts: %+v\n", workouts)
	})

	// ScheduleWorkout
	t.Run("ScheduleWorkout", func(t *testing.T) {
		// today
		today := time.Now()
		// "Afternoon Workout"

		var workout db.Workout
		er := client.Where("name = ?", "Afternoon Workout").First(&workout).Error
		// assert nil, name
		assert.Nil(t, er)
		assert.Equal(t, "Afternoon Workout", workout.Name)
		er = repo.ScheduleWorkout(&workout, today)
		// assert nil
		assert.Nil(t, er)

		workout = db.Workout{}
		er = client.Where("name = ?", "Afternoon Workout").First(&workout).Error
		// assert nil
		assert.Nil(t, er)

		// assert Date = today
		assert.Equal(t, today.Format("2006-01-02"), workout.Date.Format("2006-01-02"))
	})
	// FinishWorkout
	t.Run("FinishWorkout", func(t *testing.T) {
		// find workout Evening Workout
		var workout db.Workout
		er := client.Where("name = ?", "Evening Workout").First(&workout).Error
		// assert nil
		assert.Nil(t, er)

		err := repo.FinishWorkout(&workout)
		// assert nil
		assert.Nil(t, err)

		workout = db.Workout{}
		er = client.Where("name = ?", "Evening Workout").First(&workout).Error
		// assert nil
		assert.Nil(t, er)

		// assert Finish = true
		assert.Equal(t, true, workout.Finish)
	})

	// WorkoutToday
	t.Run("WorkoutToday", func(t *testing.T) {
		user, _ := repo.FindUserByUsername("lamdeptrai")
		// asert user.Name = "Lam"
		assert.Equal(t, "Lam", user.Name)
		today := time.Now()
		err := repo.CreateEmptyWorkout(user, "Today's Workout")
		assert.Nil(t, err)

		var workout db.Workout
		er := client.Where("name = ?", "Today's Workout").First(&workout).Error
		assert.Nil(t, er)
		// assert name = "Today's Workout"
		assert.Equal(t, "Today's Workout", workout.Name)

		er = repo.ScheduleWorkout(&workout, today)
		assert.Nil(t, er)

		workouts, err := repo.WorkoutToday(user)
		utils.PrintRed(fmt.Sprintf("workouts: %+v\n", workouts))

		assert.Nil(t, err)
		assert.Equal(t, 2, len(workouts))
	})

	// WorkoutFuture
	t.Run("WorkoutFuture", func(t *testing.T) {
		user, _ := repo.FindUserByUsername("lamdeptrai")
		future := time.Now().AddDate(0, 0, 1)
		err := repo.CreateEmptyWorkout(user, "Future Workout")
		assert.Nil(t, err)

		var workout db.Workout
		er := client.Where("name = ?", "Future Workout").First(&workout).Error
		assert.Nil(t, er)

		er = repo.ScheduleWorkout(&workout, future)
		assert.Nil(t, er)

		workouts, err := repo.WorkoutFuture(user)
		assert.Nil(t, err)
		assert.Equal(t, 2, len(workouts))
		utils.PrintRed(fmt.Sprintf("workouts: %+v\n", *workouts[0]))
		utils.PrintRed(fmt.Sprintf("workouts: %+v\n", *workouts[1]))
		// assert.Equal(t, "Future Workout", workouts[0].Name)
	})

	// WorkoutDone
	t.Run("WorkoutDone", func(t *testing.T) {
		user, _ := repo.FindUserByUsername("lamdeptrai")
		workouts, err := repo.WorkoutDone(user)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(workouts))
		assert.Equal(t, "Evening Workout", workouts[0].Name)
	})

	// PastWorkout
	t.Run("PastWorkout", func(t *testing.T) {
		user, _ := repo.FindUserByUsername("linhxinh")
		past := time.Now().AddDate(0, 0, -1)
		err := repo.CreateEmptyWorkout(user, "Past Workout")
		assert.Nil(t, err)

		var workout db.Workout
		er := client.Where("name = ?", "Past Workout").First(&workout).Error
		assert.Nil(t, er)

		er = repo.ScheduleWorkout(&workout, past)
		assert.Nil(t, er)

		workouts, err := repo.PastWorkout(user)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(workouts))
		assert.Equal(t, "Past Workout", workouts[0].Name)
	})
}
