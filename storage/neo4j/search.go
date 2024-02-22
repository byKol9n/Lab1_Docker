package neo4j

import (
	"strconv"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func (n *Neo4j) GetVisited(lectures []int) ([]int, []int, error) {
	var lessons, students []int

	var itemsInterface []interface{}
    for _, item := range lectures {
        itemsInterface = append(itemsInterface, item)
    }

	result, _ := neo4j.ExecuteQuery(n.context, n.conn,
		`MATCH (l:Lesson)-[:BELONGS_TO]->(lec:Lecture)
		WHERE lec.id_lecture IN $lectures
		MATCH (s:Schedule)
		WHERE s.id_lesson = l.id_lesson
		WITH DISTINCT s.id_group AS id_group
		MATCH (st:Student)
		WHERE st.id_group = id_group
		RETURN DISTINCT st.id_student AS s, l.id_lesson AS l`,
		map[string]interface{}{
            "lectures": itemsInterface,
        }, neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("graph"))

		for _, record := range result.Records {
			l, _ := record.Get("l")
			lid, _ := strconv.Atoi(l.(string))
			lessons = append(lessons, lid)

			s, _ := record.Get("s")
			sid, _ := strconv.Atoi(s.(string))
			students = append(students, sid)
		}

	return lessons, students, nil
}