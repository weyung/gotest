// service/service.go
package service

import (
	"github.com/weyung/gotest/model"
	"github.com/weyung/gotest/util"
)

// GetScoreByStudentID returns the score of a student by student ID
func GetScoreByStudentID(studentID string) (*model.Score, error) {
	score := &model.Score{}
	err := util.DB.Get(score, "SELECT student_id, subject, score FROM scores WHERE student_id = $1", studentID)
	if err != nil {
		return nil, err
	}
	return score, nil
}
