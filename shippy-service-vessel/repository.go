package main

import (
	"context"

	pb "github.com/galahade/shippy/shippy-service-vessel/proto/vessel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository interface {
	FindAvailable(ctx context.Context, spec *Specification) (*Vessel, error)
	Create(ctx context.Context, vessel *Vessel) error
}

//MongoRepository contain a mongo collection struct form CRUD
type MongoRepository struct {
	collection *mongo.Collection
}

//Specification is a json mapper for mongoDB
type Specification struct {
	Capacity  int32 `json:"capacity"`
	MaxWeight int32 `json:"max_weight"`
}

//MarshalSpecification -
func MarshalSpecification(spec *pb.Specification) *Specification {
	return &Specification{
		Capacity:  spec.Capacity,
		MaxWeight: spec.MaxWeight,
	}
}

//UnmarshalSpecification -
func UnmarshalSpecification(spec *Specification) *pb.Specification {
	return &pb.Specification{
		Capacity:  spec.Capacity,
		MaxWeight: spec.MaxWeight,
	}
}

//MarshalVessel -
func MarshalVessel(vessel *pb.Vessel) *Vessel {
	return &Vessel{
		ID:        vessel.Id,
		Capacity:  vessel.Capacity,
		MaxWeight: vessel.MaxWeight,
		Name:      vessel.Name,
		Available: vessel.Available,
		OwnerID:   vessel.OwnerId,
	}
}

//UnmarshalVessel -
func UnmarshalVessel(vessel *Vessel) *pb.Vessel {
	return &pb.Vessel{
		Id:        vessel.ID,
		Capacity:  vessel.Capacity,
		MaxWeight: vessel.MaxWeight,
		Name:      vessel.Name,
		Available: vessel.Available,
		OwnerId:   vessel.OwnerID,
	}
}

// Vessel -
type Vessel struct {
	ID        string
	Capacity  int32
	Name      string
	Available bool
	OwnerID   string
	MaxWeight int32
}

// FindAvailable - checks a specification against a map of vessels,
// if capacity and max weight are below a vessels capacity and max weight,
// then return that vessel.
func (repository *MongoRepository) FindAvailable(ctx context.Context, spec *Specification) (*Vessel, error) {
	filter := bson.D{{
		"capacity",
		bson.D{{
			"$lte",
			spec.Capacity,
		}, {
			"$lte",
			spec.MaxWeight,
		}},
	}}
	vessel := &Vessel{}
	if err := repository.collection.FindOne(ctx, filter).Decode(vessel); err != nil {
		return nil, err
	}
	return vessel, nil
}

// Create a new vessel
func (repository *MongoRepository) Create(ctx context.Context, vessel *Vessel) error {
	_, err := repository.collection.InsertOne(ctx, vessel)
	return err
}
