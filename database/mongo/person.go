package main

import (
	"github.com/lordking/toolbox/common"
	"github.com/lordking/toolbox/database/mongo"
	"github.com/lordking/toolbox/log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	collectionName = "person"
)

type (
	Person struct {
		db *mongo.Mongo
	}

	//Person 用户数据对象
	PersonVO struct {
		Id    bson.ObjectId `json:"id" bson:"_id"`
		Name  string        `json:"name" bson:"name"`
		Phone string        `json:"phone" bson:"phone"`
	}
)

func (p *Person) insert(obj *PersonVO) {

	obj.Id = bson.NewObjectId()
	collection, err := p.db.GetCollection(collectionName)
	err = collection.Insert(obj)
	defer common.CheckError(err)

	log.Debug("Insert result: %s", common.PrettyObject(obj))

}

func (p *Person) findAll(name string) {

	var result []PersonVO
	collection, err := p.db.GetCollection(collectionName)
	err = collection.Find(bson.M{"name": name}).All(&result)
	defer common.CheckError(err)

	log.Debug("Find result: %s", common.PrettyObject(result))

}

func (p *Person) updateAll(name string, obj *PersonVO) {

	var result *mgo.ChangeInfo
	collection, err := p.db.GetCollection(collectionName)
	result, err = collection.UpdateAll(bson.M{"name": name}, bson.M{"$set": bson.M{"phone": obj.Phone}})
	defer common.CheckError(err)

	log.Debug("Update result: %s", common.PrettyObject(result))

}

func (p *Person) removeAll(name string) {

	var result *mgo.ChangeInfo
	collection, err := p.db.GetCollection(collectionName)
	result, err = collection.RemoveAll(bson.M{"name": name})
	defer common.CheckError(err)

	log.Debug("Remove result: %s", common.PrettyObject(result))
}

func NewPerson(db *mongo.Mongo) (*Person, error) {

	err := db.Connect()

	return &Person{
		db: db,
	}, err
}
