package repository

import (
	"time"

	"github.com/aditya-suripeddi/covstats/model"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// RegionInfoRepositoryMongo is an instance connection of MongoDB
type RegionInfoRepositoryMongo struct {
	db         *mgo.Database
	collection string
}

// NewRegionInfoRepositoryMongo is function that will return userRepositoryMongo
func NewRegionInfoRepositoryMongo(db *mgo.Database, collection string) *RegionInfoRepositoryMongo {
	return &RegionInfoRepositoryMongo{
		db:         db,
		collection: collection,
	}
}

// Save is a function to create new user
func (r *RegionInfoRepositoryMongo) Save(regionInfo *model.RegionInfo) error {
	err := r.db.C(r.collection).Insert(regionInfo)
	return err
}

// Update is a function to update existing user
func (r *RegionInfoRepositoryMongo) Update(rno string, regionInfo *model.RegionInfo) error {
	regionInfo.UpdatedAt = time.Now()
	err := r.db.C(r.collection).Update(bson.M{"rno": rno}, regionInfo)
	return err
}

// Delete is a function to remove a record from User list
func (r *RegionInfoRepositoryMongo) Delete(rno string) error {
	err := r.db.C(r.collection).Remove(bson.M{"rno": rno})
	return err
}

// FindByRegion is a function to get one user by ID
func (r *RegionInfoRepositoryMongo) FindByRegion(rname string) (*model.RegionInfo, error) {
	var regionInfo model.RegionInfo
	err := r.db.C(r.collection).Find(bson.M{"region_name": rname}).One(&regionInfo)
	if err != nil {
		return nil, err
	}
	return &regionInfo, nil
}

// FindAll is a function to get User list
func (r *RegionInfoRepositoryMongo) FindAll() (model.Region, error) {
	var data model.Region
	err := r.db.C(r.collection).Find(bson.M{}).All(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
