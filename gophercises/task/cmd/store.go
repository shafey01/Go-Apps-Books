package cmd

import (
	"encoding/binary"
	"encoding/json"
	"fmt"

	"github.com/boltdb/bolt"
)

const dbFileName = "tasks.db"

var tasksBucketName = []byte("tasks")

type Task struct {
	ID        int    `json:"id"`
	Details   string `json:"details"`
	Completed bool   `json:"completed"`
}

func setupDB() (*bolt.DB, error) {
	db, err := bolt.Open(dbFileName, 0600, nil)
	if err != nil {
		return nil, fmt.Errorf("could not open db, %v", err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(tasksBucketName)
		if err != nil {
			return fmt.Errorf("could not create root bucket: %v", err)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("could not set up buckets, %v", err)
	}
	fmt.Println("DB Setup Done")
	return db, nil
}

func SetTask(db *bolt.DB, task *Task) error {
	err := db.Update(func(tx *bolt.Tx) error {
		id, err := tx.Bucket(tasksBucketName).NextSequence()
		if err != nil {
			return err

		}
		task.ID = int(id)
		taskMarshal, err := json.Marshal(&task)
		if err != nil {
			return fmt.Errorf("could not marshal Task json: %v", err)
		}
		err = tx.Bucket(tasksBucketName).Put(itob(task.ID), taskMarshal)
		if err != nil {
			return fmt.Errorf("could not set Task: %v", err)
		}
		return nil
	})
	fmt.Println("Set Task")
	return err
}

func ListTasks(db *bolt.DB, completed bool) ([]*Task, error) {
	var tasks []*Task

	return tasks, db.View(func(tx *bolt.Tx) error {

		return tx.Bucket(tasksBucketName).ForEach(func(_, b []byte) error {
			var task Task
			if err := json.Unmarshal(b, &task); err != nil {
				return err
			}
			if task.Completed != completed {
				return nil
			}
			tasks = append(tasks, &task)
			return nil
		})
	})

}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

// func CreateTask(task *Task) error {
// 	return withDB(func(db *bolt.DB) error {
// 		return db.Update(func(tx *bolt.Tx) error {
// 			bucket := tx.Bucket(tasksBucketName)

// 			id, err := bucket.NextSequence()
// 			if err != nil {
// 				return err
// 			}
// 			task.ID = int(id)

// 			b, err := json.Marshal(&task)
// 			if err != nil {
// 				return err
// 			}

// 			return bucket.Put(itob(task.ID), b)
// 		})
// 	})
// }
