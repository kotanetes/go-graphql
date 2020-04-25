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

func (r *queryResolver) Teachers(ctx context.Context) ([]*Teacher, error) {

	var (
		results  []Teacher
		teachers []*Teacher
		err      error
	)
	results, err = getTeachers()
	if err != nil {
		return nil, err
	}

	for i := range results {
		teachers = append(teachers, &results[i])
	}

	return teachers, nil

}

func getTeachers() ([]Teacher, error) {
	var (
		elem    Teacher
		results []Teacher
	)
	session := database.DbSession

	coll := session.Database("school").Collection("teacher")

	cur, err := coll.Find(context.TODO(), bson.M{})
	if err != nil {
		return results, err
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		err := cur.Decode(&elem)
		if err != nil {
			return results, err
		}
		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	// Close the cursor once finished
	cur.Close(context.TODO())
	return results, nil
}

func (r *queryResolver) Teacher(ctx context.Context, firstName string) (*Teacher, error) {
	data, err := getTeacher(firstName)
	return &data, err
}
func getTeacher(name string) (Teacher, error) {
	var (
		teacher Teacher
	)
	session := database.DbSession

	coll := session.Database("school").Collection("teacher")

	result := coll.FindOne(context.TODO(), bson.M{"firstName": name})

	err := result.Decode(&teacher)
	if err != nil {
		return teacher, err
	}
	return teacher, nil
}
