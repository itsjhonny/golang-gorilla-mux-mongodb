package configs

import "fmt"

func ConnectDB() {

	t, err := load()

	if err != nil {
		panic(err)
	}

	fmt.Printf(t.Database)

	/*
		client, err := mongo.NewClient(options.Client().ApplyURI(load().Uri))
		if err != nil {
			log.Fatal(err)
		}

		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err = client.Connect(ctx)
		if err != nil {
			log.Fatal(err)
		}

		//ping the database
		err = client.Ping(ctx, nil)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Connected to MongoDB")
		return client
	*/
}
