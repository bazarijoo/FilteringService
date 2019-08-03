package persistence

import (
	"FilteringService/model"
	"bytes"
	"encoding/json"

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
			return err
		}
		_, err = conn.Do("RPUSH", "rectangles", rectJson)
		if err != nil {
			return err
		}

	}

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
			panic(err)
		}
		rectangles = append(rectangles, rect)
	}

	return rectangles

}

// items, err := redis.ByteSlices(d.Conn.Do("LRANGE", "objects", "0", "-1"))
// if err != nil {
//    // handle error
// }
// var values []*Object
// for _, item := range items {
//     var v Object
//     if err := gob.NewDecoder(bytes.NewReader(item)).Decode(&v); err != nil {
//         // handle error
//     }
//     values = append(values, &v)
// }
