package redis

import (
	"encoding/json"
	"noname_team_project/model"
)

func (r *Redis) Set(student model.Student) error {
	marshalStudent, err := json.Marshal(student)
	if err != nil {
		return err
	}

	if err := r.conn.Set(ctx, student.Id, marshalStudent, 0).Err(); err != nil {
		return err
	}

	return nil
}

func (r *Redis) Get(studentId string) (string, error) {
	marshalStudent, err := r.conn.Get(ctx, studentId).Result()
	if err != nil {
		return err.Error(), err
	}

	return marshalStudent, nil
}
