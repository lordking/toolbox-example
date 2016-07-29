package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	ws "goutils"
)

type Blog struct {
	Id         bson.ObjectId `json:"id" bson:"_id"`
	Subject    string        `json:"subject" bson:"subject"`
	Blog       string        `json:"blog" bson:"blog"`
	Author     string        `json:"author" bson:"author"`
	CreateTime int64         `json:"createTime" bson:"createTime"`
	UpdateTime int64         `json:"updateTime" bson:"updateTime"`
}

func (b *Blog) Create() error {
	mongo := (ws.DataSourceInstance).(*ws.Mongo)
	collection, err := mongo.GetCollection("blog")

	if err != nil {
		return ws.ToError(ws.ErrCodeInternal, err)
	}

	b.Id = bson.NewObjectId() //生成id
	b.CreateTime = time.Now().Unix()
	b.UpdateTime = b.CreateTime
	err = collection.Insert(b)

	if err != nil {
		return ws.ToError(ws.ErrCodedParams, err)
	}
	return nil
}

func (b *Blog) Find(start int, number int) ([]Blog, error) {
	mongo := (ws.DataSourceInstance).(*ws.Mongo)
	collection, err := mongo.GetCollection("blog")

	if err != nil {
		return nil, ws.ToError(ws.ErrCodeInternal, err)
	}

	//测试查询
	var result []Blog
	err = collection.Find(nil).Skip(start).Limit(number).All(&result)

	if err != nil {
		return nil, ws.ToError(ws.ErrCodedParams, err)
	}
	return result, nil
}

func (b *Blog) Update(id string) error {
	mongo := (ws.DataSourceInstance).(*ws.Mongo)
	collection, err := mongo.GetCollection("blog")

	if err != nil {
		return ws.ToError(ws.ErrCodedParams, err)
	}

	//测试更新
	b.Id = bson.ObjectIdHex(id)

	b.UpdateTime = time.Now().Unix()
	data, err := bson.Marshal(b)
	if err != nil {
		return ws.ToError(ws.ErrCodedParams, err)
	}

	var updateJson map[string]interface{}
	err = bson.Unmarshal(data, &updateJson)
	if err != nil {
		return ws.ToError(ws.ErrCodedParams, err)
	}

	delete(updateJson, "_id")
	delete(updateJson, "createTime")
	for key, value := range updateJson {
		if value == "" || value == nil {
			delete(updateJson, key)
		}
	}

	err = collection.UpdateId(b.Id, bson.M{"$set": updateJson})

	if err != nil {
		return ws.ToError(ws.ErrCodedParams, err)
	}
	return nil
}

func (b *Blog) Delete(id string) error {
	mongo := (ws.DataSourceInstance).(*ws.Mongo)
	collection, err := mongo.GetCollection("blog")

	if err != nil {
		return ws.ToError(ws.ErrCodeInternal, err)
	}

	//测试查询
	objId := bson.ObjectIdHex(id)
	err = collection.RemoveId(objId)

	if err != nil {
		return ws.ToError(ws.ErrCodedParams, err)
	}
	return err
}
