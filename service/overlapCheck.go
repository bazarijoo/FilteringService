package service

import (
	"FilteringService/model"
	"time"
)

func FindOverlappedRectangles(main model.Rectangle, inputs []model.Rectangle) []model.Rectangle {

	var overlappedRectangles []model.Rectangle
	currentTime := time.Now().Local()

	for _, rect := range inputs {

		//if one rectangle is above the other or next to the other, they're not overlapped.
		aboveCondition := main.Y < rect.Y-rect.Height || rect.Y < main.Y-main.Height
		nextToCondition := rect.X-rect.Width > main.X || main.X-main.Width > rect.X

		if !(aboveCondition && nextToCondition) {
			rect.CreatedAt = currentTime
			overlappedRectangles = append(overlappedRectangles, rect)
		}
	}

	return overlappedRectangles

}
