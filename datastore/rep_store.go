package datastore

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type RepStoreServices interface {
	FindRep(string, string) int
	UpdateRep(string, string, int)
}

type RepStore struct {
	Col *mgo.Collection
}

type RepStruct struct {
	Rep int `bson:"rep"`
}

func (store *RepStore) FindRep(category, userID string) int {

	retrieved := new(RepStruct)

	err := store.Col.Find(bson.M{"_id": bson.M{"category": category, "userID": userID}}).One(retrieved)
	if err == mgo.ErrNotFound {
		store.Col.Insert(bson.M{"_id": bson.M{"category": category, "userID": userID}, "rep": 5})
	} else if err != nil {
		log.Fatal(err)
	}

	return retrieved.Rep
}

func (store *RepStore) UpdateRep(category, userID string, rep int) {

	err := store.Col.Update(bson.M{"_id": bson.M{"category": category, "userID": userID}}, bson.M{"$inc": bson.M{"rep": rep}})

	if err == mgo.ErrNotFound {
		store.Col.Insert(bson.M{"_id": bson.M{"category": category, "userID": userID}, "rep": (5 + rep)})
	} else if err != nil {
		log.Fatal(err)
	}
}
