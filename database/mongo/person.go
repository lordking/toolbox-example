package main

import (
	"github.com/lordking/toolbox/common"
	"github.com/lordking/toolbox/database"
	"github.com/lordking/toolbox/database/mongo"
	"github.com/lordking/toolbox/log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Person 用户数据对象
type Person struct {
	Id    bson.ObjectId `json:"id" bson:"_id"`
	Name  string        `json:"name" bson:"name"`
	Phone string        `json:"phone" bson:"phone"`
}

func (p *Person) insert() {
	//获取单例
	db := (database.Instance).(*mongo.Mongo)

	err := db.Connect()
	defer common.CheckFatal(err)

	p.Id = bson.NewObjectId()
	collection, err := db.GetCollection(collectionName)
	err = collection.Insert(p)
	defer common.CheckError(err)

	log.Debug("Insert result: %s", common.PrettyObject(p))

	db.Close()
}

func (p *Person) findAll(name string) {
	//获取单例
	db := (database.Instance).(*mongo.Mongo)

	err := db.Connect()
	defer common.CheckFatal(err)

	var result []Person
	collection, err := db.GetCollection(collectionName)
	err = collection.Find(bson.M{"name": name}).All(&result)
	defer common.CheckError(err)

	log.Debug("Find result: %s", common.PrettyObject(result))

	db.Close()
}

func (p *Person) updateAll(name string, phone string) {
	//获取单例
	db := (database.Instance).(*mongo.Mongo)

	err := db.Connect()
	defer common.CheckFatal(err)

	var result *mgo.ChangeInfo
	collection, err := db.GetCollection(collectionName)
	result, err = collection.UpdateAll(bson.M{"name": name}, bson.M{"$set": bson.M{"phone": phone}})
	defer common.CheckError(err)

	log.Debug("Update result: %s", common.PrettyObject(result))

	db.Close()
}

func (p *Person) removeAll(name string) {
	//获取单例
	db := (database.Instance).(*mongo.Mongo)

	err := db.Connect()
	defer common.CheckFatal(err)

	var result *mgo.ChangeInfo
	collection, err := db.GetCollection(collectionName)
	result, err = collection.RemoveAll(bson.M{"name": name})
	defer common.CheckError(err)

	log.Debug("Remove result: %s", common.PrettyObject(result))

	db.Close()
}
