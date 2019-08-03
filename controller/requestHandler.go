package controller

import (
	"FilteringService/model"
	"FilteringService/model/requestBody"
	"FilteringService/persistence"
	"FilteringService/service"

	"encoding/json"
	"net/http"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {

	var postRequest requestBody.PostRequest

	err := json.NewDecoder(r.Body).Decode(&postRequest)
	if err != nil {
		panic(err)
	}

	var mainRectangle model.Rectangle
	var inputs []model.Rectangle

	mainRectangle = postRequest.Main
	inputs = postRequest.Input

	results := service.FindOverlappedRectangles(mainRectangle, inputs)

	persistence.InsertRecords(results)

}

func GetHandler(w http.ResponseWriter, r *http.Request) {

	results := persistence.GetRecords()

	//Marshal or convert user object back to json and write to response
	rectangles, err := json.Marshal(results)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(rectangles)

}
