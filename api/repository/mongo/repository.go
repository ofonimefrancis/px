package mongo

import (
	"context"
	"time"

	"github.com/ofonimefrancis/pixels/api/models"
	"github.com/ofonimefrancis/pixels/api/repository"
	"github.com/pkg/errors"
	mongo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoRepository struct {
	client   *mongo.Client
	database string
	timeout  time.Duration
}

func newUserMongoClient(mongoURI string, connectionTimeout int) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(connectionTimeout)*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}
	defer cancel()

	return client, err
}

func NewUserRepository(mongoURI, mongoDB string, connectionTimeout int) (repository.UserRepository, error) {
	datastore := &userMongoRepository{
		timeout:     time.Duration(connectionTimeout) * time.Second,
		databaseURI: mongoDB,
	}

	client, err := newUserMongoClient(mongoURI, connectionTimeout)
	if err != nil {
		return nil, errors.Wrap(err, "datastore.NewUserMongoRepository")
	}
	datastore.client = client
	return datastore, nil
}

func (r *mongoRepository) Find(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	user := &models.User{}
	collection := r.client.Database(r.database).Collection("users")
	filter := bson.M{"email": email}
	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err === mongo.ErrNoDocuments {
			return nil, errors.Wrap(errors.New("User not found"), "repository.User.Find")
		}
		return nil, errors.Wrap(err, "repository.User.Find")
	}
	return user, nil
}

func (r *mongoRepository) Store(m *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	collection := r.client.Database(r.database).Collection("users")
	_, err := collection.InsertOne(
		ctx, 
		bson.M{
			"first_name": m.FirstName,
			"last_name": m.LastName,
			"email": m.Email,
		},
	)
	if err != nil {
		return errors.Wrap(err, "repository.User.Store")
	}
	return nil
}
