package repositories

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"teslaUsers/src/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	collection = "users"
)

type userRepoMongo struct {
	mongoClient *mongo.Client
	database    string
}

func NewUserMongoRepo(mongoClient *mongo.Client, database string) *userRepoMongo {
	return &userRepoMongo{
		mongoClient: mongoClient,
		database:    database,
	}
}

func (userRepo *userRepoMongo) AddUser(user models.User) (models.User, error) {
	user.ID = primitive.NewObjectID()
	_, err := userRepo.mongoClient.Database(userRepo.database).Collection(collection).InsertOne(context.TODO(), user)
	return user, err
}

func (userRepo *userRepoMongo) AddPaymentMethod(mail string, paymentMethod string) error {
	_, err := userRepo.mongoClient.Database(userRepo.database).Collection(collection).UpdateOne(context.TODO(), bson.M{"email": mail}, bson.M{"$push": bson.M{"payment_methods": paymentMethod}})
	return err
}

func (userRepo *userRepoMongo) GetUsers() ([]models.User, error) {
	cursorUsers, err := userRepo.mongoClient.Database(userRepo.database).Collection(collection).Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}
	var users []models.User
	err = cursorUsers.All(context.TODO(), &users)

	return users, err
}

func (userRepo *userRepoMongo) GetUser(mail string) (models.User, error) {
	var user models.User
	query := bson.M{"email": mail}
	err := userRepo.mongoClient.Database(userRepo.database).
		Collection(collection).FindOne(context.TODO(), query).Decode(&user)
	return user, err
}
