package query

import "strings"

type ImageQuery struct {
	ids []uint64
}

func NewImageQuery() *ImageQuery {
	return &ImageQuery{}
}

func (q *ImageQuery) ID(
	ID uint64,
) *ImageQuery {
	q.ids = append(q.ids, ID)

	return q
}

func (q *ImageQuery) Build() (wheres string, args []interface{}) {
	whereCols := make([]string, 0)

	if len(q.ids) != 0 {
		inCondition := "id IN (?"

		for i := 1; i < len(q.ids); i++ {
			inCondition = inCondition + ",?"
		}
		inCondition = inCondition + ")"
		whereCols = append(whereCols, inCondition)

		for _, id := range q.ids {
			args = append(args, id)
		}
	}

	wheres = strings.Join(whereCols, " AND ")

	return wheres, args
}
