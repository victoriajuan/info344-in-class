package tasks

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type completedFilter struct {
	Completed bool
}

type updateDoc struct {
	Completed  bool
	ModifiedAt time.Time
}

//MongoStore implements Store for MongoDB
type MongoStore struct {
	session *mgo.Session
	dbname  string
	colname string
	col     *mgo.Collection
}

//NewMongoStore constructs a new MongoStore
func NewMongoStore(sess *mgo.Session, dbName string, collectionName string) *MongoStore {
	if sess == nil {
		panic("nil pointer passed for session")
	}
	return &MongoStore{
		session: sess,
		dbname:  dbName,
		colname: collectionName,
		col:     sess.DB(dbName).C(collectionName),
	}
}

//Insert inserts a new task into the store
func (s *MongoStore) Insert(nt *NewTask) (*Task, error) {
	task, err := nt.ToTask()
	if err != nil {
		return nil, err
	}
	if err := s.col.Insert(task); err != nil {
		return nil, fmt.Errorf("error inserting task: %v", err)
	}
	return task, nil
}

//GetAll gets all tasks (up to AllTasksLimit) with a particular `completed` value
func (s *MongoStore) GetAll(completed bool) ([]*Task, error) {
	tasks := []*Task{}
	filter := &completedFilter{completed}
	if err := s.col.Find(filter).Limit(AllTasksLimit).All(&tasks); err != nil {
		return nil, fmt.Errorf("error getting tasks: %v", err)
	}
	return tasks, nil
}

//Update updates the task with `id` based on the values in `tu`
func (s *MongoStore) Update(id bson.ObjectId, tu *TaskUpdates) (*Task, error) {
	upd := &updateDoc{
		Completed:  tu.Completed,
		ModifiedAt: time.Now(),
	}
	change := mgo.Change{
		Update:    bson.M{"$set": upd},
		ReturnNew: true,
	}
	task := &Task{}
	if _, err := s.col.FindId(id).Apply(change, task); err != nil {
		return nil, fmt.Errorf("error updating task: %v", err)
	}
	return task, nil
}
