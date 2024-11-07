package mongo

import (
	"context"
	"time"

	"github.com/dotpep/golang-ddd-architecture/aggregate"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	db       *mongo.Database
	customer *mongo.Collection
}

// mongoCustomer is a internal type
// that is used to store a CustomerAggregate inside this Repository.
// We have to use an internal struct to avoid coupling,
// this implementation shouldn't have any coupling to Aggregate.
// bson tags is mongodb.
type mongoCustomer struct {
	ID   uuid.UUID `bson:"id"`
	Name string    `bson:"name"`
}

// One very common approach is to have some sort of converters/formats between the formats,
// since we're not operating on Aggregates straight away inside this,
// we can sort to have these helper functions that converts between formats for us.
// so now we can easily go from Aggregate to a local internal struct, which we will use,
// and it is same in other way around for mongoCustomer to Aggregate.Customer

// NewFromCustomer is converter of Aggregate.Customer to MongoCustomer
func NewFromCustomer(c aggregate.Customer) mongoCustomer {
	return mongoCustomer{
		ID:   c.GetID(),
		Name: c.GetName(),
	}
}

// mongoCustomer ToAggregate function converter and formatter
func (m mongoCustomer) ToAggregate() aggregate.Customer {
	c := aggregate.Customer{}

	c.SetID(m.ID)
	c.SetName(m.Name)

	return c
}

// New is function for connecting to MongoDB
func New(ctx context.Context, connectionString string) (*MongoRepository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, err
	}

	db := client.Database("golang-ddd")
	customers := db.Collection("customers")

	return &MongoRepository{
		db:       db,
		customer: customers,
	}, nil
}

// Get(uuid.UUID) (aggregate.Customer, error)
func (mr *MongoRepository) Get(id uuid.UUID) (aggregate.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result := mr.customer.FindOne(ctx, bson.M{"id": id})

	var c mongoCustomer

	err := result.Decode(&c)
	if err != nil {
		return aggregate.Customer{}, err
	}

	return c.ToAggregate(), nil
}

func (mr *MongoRepository) Add(c aggregate.Customer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	internal := NewFromCustomer(c)

	_, err := mr.customer.InsertOne(ctx, internal)
	if err != nil {
		return nil
	}

	return nil
}

func (mr *MongoRepository) Update(c aggregate.Customer) error {
	panic("Not Implemented Yet!")
}
