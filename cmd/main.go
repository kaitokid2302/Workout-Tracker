package main

import "github.com/kaitokid2302/Workout-Tracker/internal/app"

func main() {
	app := app.App{}
	app.SetUp()
	app.Run()
}
