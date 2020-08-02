package user

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoRepository struct {
	db *mongo.Client
}

const databaseName = "user"

// NewMongo methods
func NewMongo(url string) (Repository, error) {
	credential := options.Credential{
		Username: "argadev",
		Password: "123456",
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url).SetAuth(credential))
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return &mongoRepository{client}, nil
}

func (r *mongoRepository) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	defer func() {
		if err := r.db.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func (r *mongoRepository) AddUser(ctx context.Context, name string, email string, address string, roleID uint32, genderID uint32, bloodID uint32, birthOfDate string, contact string) error {
	user := User{
		Name:        name,
		Email:       email,
		Address:     address,
		RoleID:      roleID,
		GenderID:    genderID,
		BloodID:     bloodID,
		BirthOfDate: birthOfDate,
		Contact:     contact,
	}
	result, err := r.db.Database(databaseName).Collection("users").InsertOne(ctx, user)
	if err != nil {
		return err
	}
	fmt.Printf("Insert One User => %s\n", result.InsertedID)
	return nil

}

func (r *mongoRepository) GetUsers(ctx context.Context) ([]User, error) {
	var users []User

	cursor, err := r.db.Database(databaseName).Collection("users").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user User
		_ = cursor.Decode(&user)
		users = append(users, user)
	}
	if err = cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *mongoRepository) GetUserByID(ctx context.Context, id string) (*User, error) {

	var user User
	userID, _ := primitive.ObjectIDFromHex(id)

	err := r.db.Database(databaseName).Collection("users").FindOne(ctx, User{ID: userID}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
