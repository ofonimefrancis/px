package mongodb

import (
	"context"
	"time"

	commons "github.com/ofonimefrancis/pixels/pkg/commons/error"
	"github.com/ofonimefrancis/pixels/pkg/datastore"
	"github.com/ofonimefrancis/pixels/pkg/models"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type userMongoRepository struct {
	client      *mongo.Client
	databaseURI string
	timeout     time.Duration
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

func NewUserRepository(mongoURI, mongoDB string, connectionTimeout int) (datastore.UserRepository, error) {
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

func (datastore *userMongoRepository) GetAll() ([]models.User, error) {
	var users []models.User
	ctx, cancel := context.WithTimeout(context.Background(), datastore.timeout)
	defer cancel()
	collection := datastore.client.Database(datastore.databaseURI).Collection(new(models.User).TableName())
	cursor, err := collection.Find(ctx, nil)
	if err != nil {
		return users, err
	}
	if err := cursor.Decode(&users); err != nil {
		return users, err
	}
	return users, nil
}

func (datastore *userMongoRepository) GetBy(filter map[string]interface{}) (*models.User, error) {
	user := new(models.User)
	ctx, cancel := context.WithTimeout(context.Background(), datastore.timeout)
	defer cancel()

	collection := datastore.client.Database(datastore.databaseURI).Collection(user.TableName())
	err := collection.FindOne(ctx, filter).Decode(user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return user, errors.Wrap(commons.ErrUserNotFound, "datastore.User.GetById")
		}
		return user, errors.Wrap(err, "datastore.User.GetById")
	}
	return user, nil
}

func (datastore *userMongoRepository) Store(data *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), datastore.timeout)
	defer cancel()

	collection := datastore.client.Database(datastore.databaseURI).Collection(data.TableName())
	if _, err := collection.InsertOne(ctx, data); err != nil {
		return errors.Wrap(err, "datastore.User.Store")
	}

	return nil
}

func (datastore *userMongoRepository) Update(data map[string]interface{}, id string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), datastore.timeout)
	defer cancel()

	collection := datastore.client.Database(datastore.databaseURI).Collection(new(models.User).TableName())

	user := new(models.User)

	filter := map[string]interface{}{"id": id}

	if res, err := collection.UpdateOne(ctx, filter, map[string]interface{}{"$set": data}, options.Update().SetUpsert(false)); err != nil {
		return user, errors.Wrap(err, "datastore.User.Update")
	} else {
		if res.MatchedCount == 0 && res.MatchedCount == 0 {
			return user, errors.Wrap(commons.ErrUserNotFound, "datastore.User.Update")
		}
	}

	user, err := datastore.GetBy(filter)
	if err != nil {
		return user, errors.Wrap(err, "datastore.User.Update")
	}

	return user, nil

}

func (datastore *userMongoRepository) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), datastore.timeout)
	defer cancel()

	filter := map[string]interface{}{"id": id}

	collection := datastore.client.Database(datastore.databaseURI).Collection(new(models.User).TableName())
	if res, err := collection.DeleteOne(ctx, filter); err != nil {
		return errors.Wrap(err, "datastore.User.Delete")
	} else {
		if res.DeletedCount == 0 {
			return errors.Wrap(commons.ErrUserNotFound, "datastore.User.Delere")
		}
	}
	return nil
}

func (datastore *userMongoRepository) Authenticate(email, password string) (bool, *models.User, error) {
	user := new(models.User)

	ctx, cancel := context.WithTimeout(context.Background(), datastore.timeout)
	defer cancel()

	collection := datastore.client.Database(datastore.databaseURI).Collection(user.TableName())
	findUserWithEmailFilter := map[string]interface{}{"email": email}
	if err := collection.FindOne(ctx, findUserWithEmailFilter).Decode(user); err != nil {
		if err == mongo.ErrNoDocuments {
			return false, user, errors.Wrap(commons.ErrUserNotFound, "datastore.User.Authenticate")
		}
		return false, user, errors.Wrap(err, "datastore.User.Authenticate")
	}

	//TODO: Check if password match that in the database
	// and return accordingly

	return true, user, nil
}
