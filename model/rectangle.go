package model

import "time"

type Rectangle struct {
	X         int
	Y         int
	Width     int
	Height    int
	CreatedAt time.Time
}

func FindOverlappedRectangles(main Rectangle, inputs []Rectangle) []Rectangle {

	var overlappedRectangles []Rectangle
	currentTime := time.Now().Local()

	for _, rect := range inputs {

		//if one rectangle is above the other or next to the other, they're not overlapped.
		above := main.Y < rect.Y-rect.Height || rect.Y < main.Y-main.Height
		nextTo := rect.X-rect.Width > main.X || main.X-main.Width > rect.X

		if !(above && nextTo) {
			rect.CreatedAt = currentTime
			overlappedRectangles = append(overlappedRectangles, rect)
		}
	}

	return overlappedRectangles

}
