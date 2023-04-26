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

	// Define the data to be inserted
	data := []string{"key1", "value1", "key2", "value2", "key3", "value3"}

	// Use pipeline to insert data in bulk
	pipe := client.Pipeline()
	for i := 0; i < len(data); i += 2 {
		pipe.Set(data[i], data[i+1], 0)
	}
	_, err = pipe.Exec()
	if err != nil {
		panic(err)
	}

	// Retrieve the data from Redis
	for i := 0; i < len(data); i += 2 {
		result, err := client.Get(data[i]).Result()
		if err != nil {
			panic(err)
		}
		fmt.Println("Key:", data[i], "Value:", result)
	}
}

/*
In this example, we create a Redis client using the "github.com/go-redis/redis" package and 
connect to a Redis instance running on the local machine. We then define an array of strings 
representing the data to be inserted, where the even indices represent the keys and the odd 
indices represent the values.

We use the pipeline feature of the Redis client to insert the data in bulk. We iterate over 
the array using a for loop and use the Set() method to set the keys and values in Redis. 
The 0 argument is the expiration time, which we set to 0 to indicate that the data should not expire.

Finally, we retrieve the data from Redis using the Get() method and print the results.
*/
