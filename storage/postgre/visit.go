package postgre

import (
	"noname_team_project/model"
	"github.com/lib/pq"
)

func (p *Postgre) GetVisitRate(studentArray, lessonsArray []int) ([]model.Rate, error) {
	var rates []model.Rate
	
	rows, err := p.conn.Query("SELECT students.id, (SELECT count(id) FROM attendances WHERE stud_id = students.id AND sched_id IN (SELECT id FROM schedules WHERE lesson_id = ANY($1)))/(SELECT count(id) FROM lessons WHERE id = ANY($1))::float AS score FROM students WHERE id = ANY($2) ORDER BY score ASC LIMIT 10;", pq.Array(lessonsArray), pq.Array(studentArray))
	if err != nil {
		return rates, err
	}
	defer rows.Close()

	for rows.Next(){
        rate := model.Rate{}
        err := rows.Scan(&rate.Id, &rate.Score)
        if err != nil{
            return rates, err
        }
        rates = append(rates, rate)
    }
	return rates, nil
}
