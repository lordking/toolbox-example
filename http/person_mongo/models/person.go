package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	ws "goutils"
)

//Person 用户数据对象
type Person struct {
	Id    bson.ObjectId `json:"id" bson:"_id"`
	Name  string        `json:"name" bson:"name"`
	Phone string        `json:"phone" bson:"phone"`
}

//Create 创建用户
func (p *Person) Create() error {
	mongo := (ws.DataSourceInstance).(*ws.Mongo)
	personCollection, err := mongo.GetCollection("person")

	if err != nil {
		return ws.ToError(ws.ErrCodeInternal, err)
	}

	//创建
	p.Id = bson.NewObjectId() //生成id
	err = personCollection.Insert(p)

	if err != nil {
		return ws.ToError(ws.ErrCodedParams, err)
	}

	return nil
}

//Find 查找用户
func (p *Person) Find(name string) ([]Person, error) {

	mongo := (ws.DataSourceInstance).(*ws.Mongo)
	personCollection, err := mongo.GetCollection("person")

	if err != nil {
		return nil, ws.ToError(ws.ErrCodeInternal, err)

	}

	//查询
	var result []Person
	err = personCollection.Find(bson.M{"name": name}).All(&result)

	if err != nil {
		return nil, ws.ToError(ws.ErrCodedParams, err)
	}

	return result, nil
}

//Update 更新用户
func (p *Person) Update(name string) (*mgo.ChangeInfo, error) {

	mongo := (ws.DataSourceInstance).(*ws.Mongo)
	personCollection, err := mongo.GetCollection("person")

	if err != nil {
		return nil, ws.ToError(ws.ErrCodeInternal, err)
	}

	//修改
	result, err := personCollection.UpdateAll(bson.M{"name": name}, bson.M{"$set": bson.M{"phone": p.Phone}})
	if err != nil {
		return nil, ws.ToError(ws.ErrCodedParams, err)
	}

	return result, nil
}

//Delete 删除用户
func (p *Person) Delete(name string) (*mgo.ChangeInfo, error) {

	mongo := (ws.DataSourceInstance).(*ws.Mongo)
	personCollection, err := mongo.GetCollection("person")

	if err != nil {
		return nil, ws.ToError(ws.ErrCodeInternal, err)

	}

	//测试查询
	result, err := personCollection.RemoveAll(bson.M{"name": name})
	if err != nil {
		return nil, ws.ToError(ws.ErrCodedParams, err)
	}

	return result, nil
}
