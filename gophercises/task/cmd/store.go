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

func MarkTaskAsCompleted(db *bolt.DB, task *Task) error {
	err := db.Update(func(tx *bolt.Tx) error {
		t := tx.Bucket(tasksBucketName).Get(itob(task.ID))
		if len(t) == 0 {
			return fmt.Errorf("There is no task with this id: %v", task.ID)
		}
		if err := json.Unmarshal(t, task); err != nil {
			exitf("%v", err)
		}
		task.Completed = true
		taskMarshal, err := json.Marshal(task)
		if err != nil {
			exitf("%v", err)
		}
		err = tx.Bucket(tasksBucketName).Put(itob(task.ID), taskMarshal)
		if err != nil {
			return fmt.Errorf("could not mark Task: %v", err)
		}
		return nil
	})
	fmt.Println("Marked Task")
	return err
}

func Delete(db *bolt.DB, task *Task) error {

	return db.Update(func(tx *bolt.Tx) error {
		t := tx.Bucket(tasksBucketName).Get(itob(task.ID))
		if len(t) == 0 {
			return fmt.Errorf("Task not found with ID=%d ", task.ID)
		}
		// if err := json.Unmarshal(t, task); err != nil {
		// 	exitf("%v", err)
		// }
		return tx.Bucket(tasksBucketName).Delete(itob(task.ID))

	})

}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
