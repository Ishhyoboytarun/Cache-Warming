package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	// Create a Redis client
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// Ping Redis to check if it's running
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}

	// Define the data to be cached
	data := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}

	// Preload Redis with the data
	for key, value := range data {
		err = client.Set(key, value, 0).Err()
		if err != nil {
			panic(err)
		}
	}

	// Retrieve the data from Redis
	for key, _ := range data {
		result, err := client.Get(key).Result()
		if err != nil {
			panic(err)
		}
		fmt.Println("Key:", key, "Value:", result)
	}
}

/*

In this example, we create a Redis client using the "github.com/go-redis/redis" package and connect 
to a Redis instance running on the local machine. We then define a map of key-value pairs to be cached 
and use a for loop to preload Redis with the data. Finally, we retrieve the data from Redis using 
the Get() method and print the results.

Note that this is a simple example and may not be suitable for large datasets. In production, you 
may need to optimize the code and use Redis pipelining or other techniques to improve performance.

*/
