// model/model.go
package model

// Score struct of score
type Score struct {
	StudentID  string `db:"student_id" json:"student_id"`
	Subject   string `db:"subject" json:"subject"`
    Score     int    `db:"score" json:"score"`
}
