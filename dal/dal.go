package dal

// Data access layer

import (
	"github.com/pkg/errors"
	"github.com/prometheus/common/log"
	"gopkg.in/mgo.v2"
)

var (
	session *mgo.Session
)

// create a global session and reuse it
func init() {
	var err error
	session, err = createSession()
	if err != nil {
		err = errors.Wrap(err, "failed to connect to DB")
		log.Fatal(err)
	}
}

// creates a new session document
func createSession() (*mgo.Session, error) {
	// connect to mongodb
	return mgo.Dial("127.0.0.1")
}

// FindOne find a single object using the selector
func FindOne(dbName, collName string, selector, data interface{}) (err error) {
	session.Refresh()
	c := session.DB(dbName).C(collName)

	if err = c.Find(selector).One(data); err != nil {
		return errors.Wrapf(err, "dal.FindOne failed, selector:%+v", selector)
	}

	return
}

// UpsertOne updates/create a single document
func UpsertOne(dbName, collName string, selector, data interface{}) (err error) {
	session.Refresh()
	c := session.DB(dbName).C(collName)

	if _, err = c.Upsert(selector, data); err != nil {
		return errors.Wrapf(err, "dal.UpsertOne failed, selector:%+v, data:%+v", selector, data)
	}

	return
}

// CreateOne create a single document
func CreateOne(dbName, collName string, data interface{}) (err error) {
	session.Refresh()
	c := session.DB(dbName).C(collName)

	if err = c.Insert(data); err != nil {
		return errors.Wrapf(err, "dal.createOne failed, data:%+v", data)
	}

	return
}

// DeleteOne deletes a single document
func DeleteOne(dbName, collName string, selector interface{}) (err error) {
	session.Refresh()
	c := session.DB(dbName).C(collName)

	if err = c.Remove(selector); err != nil {
		return errors.Wrapf(err, "dal.deleteOne failed, selector:%+v", selector)
	}

	return
}
