package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	ws "goutils"
)

type Token struct {
	Id         bson.ObjectId `json:"id" bson:"_id"`
	Token      string        `json:"token" bson:"token"`
	ExpireTime int64         `json:"expireTime" bson:"expireTime"`
	CreateTime int64         `json:"createTime" bson:"createTime"`
	UpdateTime int64         `json:"updateTime" bson:"updateTime"`
}

func (t *Token) Create() (*Token, error) {
	mongo := (ws.DataSourceInstance).(*ws.Mongo)
	collection, err := mongo.GetCollection("token")

	if err != nil {
		return nil, ws.ToError(ws.ErrCodeInternal, err)
	}

	t.Id = bson.NewObjectId() //生成id

	t.CreateTime = time.Now().Unix()
	t.UpdateTime = t.CreateTime
	t.ExpireTime = t.CreateTime + 3600
	t.Token = t.Id.Hex()

	err = collection.Insert(t)

	if err != nil {
		return nil, ws.ToError(ws.ErrCodedParams, err)
	}

	return t, nil
}

func (t *Token) Find(token string) (*Token, error) {
	mongo := (ws.DataSourceInstance).(*ws.Mongo)
	collection, err := mongo.GetCollection("token")

	if err != nil {
		return nil, ws.ToError(ws.ErrCodeInternal, err)
	}

	var result *Token
	objId := bson.ObjectIdHex(token)
	err = collection.Find(bson.M{"_id": objId}).One(&result)

	if err != nil {
		return nil, ws.ToError(ws.ErrCodedParams, err)
	}

	return result, nil
}

func (t *Token) Delete(id string) error {
	mongo := (ws.DataSourceInstance).(*ws.Mongo)
	collection, err := mongo.GetCollection("token")

	if err != nil {
		return ws.ToError(ws.ErrCodeInternal, err)
	}

	objId := bson.ObjectIdHex(id)
	err = collection.RemoveId(objId)

	if err != nil {
		return ws.ToError(ws.ErrCodedParams, err)
	}
	return err
}

/**
 * 清除过期令牌
 * @param  {[type]} t *Token        [description]
 * @return {[type]}   [description]
 */
func (t *Token) ClearExpireTokens() error {
	mongo := (ws.DataSourceInstance).(*ws.Mongo)
	collection, err := mongo.GetCollection("token")

	if err != nil {
		return ws.ToError(ws.ErrCodeInternal, err)
	}

	nowTime := time.Now().Unix()

	_, err = collection.RemoveAll(bson.M{"expireTime": bson.M{"$lt": nowTime}})

	if err != nil {
		return ws.ToError(ws.ErrCodedParams, err)
	}
	return err
}
