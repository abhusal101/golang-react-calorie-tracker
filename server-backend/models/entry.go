// golang-react-calorie-tracker
package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// this is the place where all the dish things will be entered
type Entry struct {
	ID          primitive.ObjectID `bson:"id"`
	Dish        *string            `json:"dish"`
	Fat         *float64           `json:"fat"`
	Ingredients *string            `json:"ingredients"`
	Calories    *string            `json:"calories"`
}
