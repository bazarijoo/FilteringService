package persistence

import (
	"FilteringService/model"
	"encoding/json"

	// "log"

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

	c := Pool.Get()
	defer c.Close()

	for _, rect := range rectangles {

		rectJSON, err := json.Marshal(rect)
		if err != nil {
			return err
		}

		_, err := c.Do("RPUSH", "rectangles", rectJSON)
		if err != nil {
			return err
		}

	}

	// log.Println("inserted successfully")

	return nil

}

// func GetRecords() []model.Rectangle {

// 	connection := Pool.Get()
// 	defer connection.Close()

// 	rectJSONs, err := redis.ByteSlices(connection.Do("LRANGE" "rectangles", 0, -1))

// 	if err != nil {
// 		log.Fatal("Error in fetch from database.")
// 		}

// 	rectangles := json.Unmarshal(rectJSONs)

// 	return rectangles

// }
