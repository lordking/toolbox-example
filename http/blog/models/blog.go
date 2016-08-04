package models

import (
	"time"

	"github.com/lordking/toolbox/common"
	"github.com/lordking/toolbox/database"
	"github.com/lordking/toolbox/database/mongo"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	Blog struct {
		collection *mgo.Collection
	}

	BlogVO struct {
		Id         bson.ObjectId `json:"id" bson:"_id"`
		Subject    string        `json:"subject" bson:"subject"`
		Blog       string        `json:"blog" bson:"blog"`
		Author     string        `json:"author" bson:"author"`
		createTime int64         `bson:"createTime"`
		updateTime int64         `bson:"updateTime"`
	}
)

func (b *Blog) Create(obj *BlogVO) error {

	obj.Id = bson.NewObjectId() //生成id
	obj.createTime = time.Now().Unix()
	obj.updateTime = obj.createTime

	err := b.collection.Insert(obj)
	if err != nil {
		return common.NewErrorWithOther(common.ErrCodeInternal, err)
	}

	return nil
}

func (b *Blog) Find(start int, number int) ([]BlogVO, error) {

	//测试查询
	var result []BlogVO
	err := b.collection.Find(nil).Skip(start).Limit(number).All(&result)
	if err != nil {
		return nil, common.NewErrorWithOther(common.ErrCodeInternal, err)
	}

	return result, nil
}

func (b *Blog) Update(id string, obj *BlogVO) error {

	obj.updateTime = time.Now().Unix()

	updateJson, err := mongo.UpdateJsonWith(obj)
	if err != nil {
		return common.NewErrorWithOther(common.ErrCodeInternal, err)
	}

	delete(updateJson, "createTime")

	objId := bson.ObjectIdHex(id)
	err = b.collection.UpdateId(objId, bson.M{"$set": updateJson})
	if err != nil {
		return common.NewErrorWithOther(common.ErrCodeInternal, err)
	}

	return nil
}

func (b *Blog) Delete(id string) error {

	objId := bson.ObjectIdHex(id)
	err := b.collection.RemoveId(objId)

	if err != nil {
		return common.NewErrorWithOther(common.ErrCodeInternal, err)
	}

	return err
}

func NewBlog() (*Blog, error) {

	//获取单例
	db := (database.Instance).(*mongo.Mongo)
	err := db.Connect()
	if err != nil {
		err = common.NewErrorWithOther(common.ErrCodeInternal, err)
	}

	collection, err := db.GetCollection("blog")
	if err != nil {
		err = common.NewErrorWithOther(common.ErrCodeInternal, err)
	}

	return &Blog{collection: collection}, err
}
