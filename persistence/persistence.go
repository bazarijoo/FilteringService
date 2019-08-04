package persistence

import (
	"FilteringService/model"
	"bytes"
	"encoding/json"
	"log"

	"github.com/garyburd/redigo/redis"
)

var Pool *redis.Pool

func NewPool() *redis.Pool {
	return &redis.Pool{
		// Maximum number of idle connections in the pool.
		MaxIdle: 5,
		// max number of connections
		MaxActive: 100,
		// Dial is an application supplied function for creating and
		// configuring a connection.
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

func InsertRecords(rectangles []model.Rectangle) error {

	conn := Pool.Get()
	defer conn.Close()

	for _, rect := range rectangles {

		rectJson, err := json.Marshal(rect)
		if err != nil {
			log.Fatal("Error occured in marshalling data.")
			return err
		}
		_, err = conn.Do("RPUSH", "rectangles", rectJson)
		if err != nil {
			log.Fatal("Error in inserting to database.")
			return err
		}

	}

	log.Println("successfully added data.")

	return nil

}

func GetRecords() []model.Rectangle {

	conn := Pool.Get()
	defer conn.Close()

	rectJSONs, err := redis.ByteSlices(conn.Do("LRANGE", "rectangles", 0, -1))

	if err != nil {
		panic(err)
	}

	var rectangles []model.Rectangle

	for _, item := range rectJSONs {
		var rect model.Rectangle
		err := json.NewDecoder(bytes.NewReader(item)).Decode(&rect)
		if err != nil {
			log.Fatal("Error occured in converting to objects.")
		}
		rectangles = append(rectangles, rect)
	}

	return rectangles

}
