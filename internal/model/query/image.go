package query

import "strings"

type ImageQuery struct {
	ids      []uint64
	cursor   *uint64
	pageSize *uint32
	orderBy  *string
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

func (q *ImageQuery) PageSize(
	pageSize *uint32,
) *ImageQuery {
	q.pageSize = pageSize

	return q
}

func (q *ImageQuery) Cursor(
	cursor *uint64,
) *ImageQuery {
	q.cursor = cursor

	return q
}

func (q *ImageQuery) OrderBy(
	orderBy *string,
) *ImageQuery {
	q.orderBy = orderBy

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

	if q.cursor != nil {
		inCondition := "created_at <= ?"

		whereCols = append(whereCols, inCondition)
		args = append(args, *q.cursor)
	}

	wheres = strings.Join(whereCols, " AND ")

	if q.orderBy != nil {
		wheres += " ORDER BY " + *q.orderBy
	}

	if q.pageSize != nil {
		inCondition := " LIMIT ?"

		wheres += inCondition
		args = append(args, *q.pageSize)
	}

	return wheres, args
}
