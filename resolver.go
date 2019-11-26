package go_graphql

import (
	"context"

	"github.com/go-graphql/database"
	"go.mongodb.org/mongo-driver/bson"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Teacher(ctx context.Context) ([]*Teacher, error) {

	var (
		elem     Teacher
		results  []Teacher
		teachers []*Teacher
	)
	session := database.DbSession

	coll := session.Database("school").Collection("teacher")

	cur, err := coll.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	for i := range results {
		teachers = append(teachers, &results[i])
	}

	return teachers, nil

}
