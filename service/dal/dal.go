package dal

// Data access layer

import (
	"target/service/models"

	"github.com/pkg/errors"
	"github.com/prometheus/common/log"
	"gopkg.in/mgo.v2"
)

var (
	session *mgo.Session
)

// create a global session and reuse it.
func init() {
	var err error
	session, err = createSession()
	if err != nil {
		err = errors.Errorf("failed to connect to DB, err:%v", err)
		log.Fatal(err)
	}
}

// creates a new session object
func createSession() (*mgo.Session, error) {
	// connect to mongodb
	return mgo.Dial("127.0.0.1")
}

func FindOne(dbName, collName string, selector, data interface{}) (prod *models.Product, err error) {
	prod = new(models.Product)

	session.Refresh()
	c := session.DB(dbName).C(collName)

	if err = c.Find(selector).One(data); err != nil {
		return nil, errors.Errorf("findOne failed, selector:%+v, err:%v", selector, err)
	}

	return
}

func UpdateOne(dbName, collName string, selector, data interface{}) (err error) {
	session.Refresh()
	c := session.DB(dbName).C(collName)

	if err = c.Update(selector, data); err != nil {
		return errors.Errorf("updateOne failed, selector:%+v, data:%+v, err:%v", selector, data, err)
	}

	return
}

func CreateOne(dbName, collName string, data interface{}) (err error) {
	session.Refresh()
	c := session.DB(dbName).C(collName)

	if err = c.Insert(data); err != nil {
		return errors.Errorf("createOne failed, data:%+v, err:%v", data, err)
	}

	return
}

func DeleteOne(dbName, collName string, selector interface{}) (err error) {
	session.Refresh()
	c := session.DB(dbName).C(collName)

	if err = c.Remove(selector); err != nil {
		return errors.Errorf("deleteOne failed, selector:%+v, err:%v", selector, err)
	}

	return
}
