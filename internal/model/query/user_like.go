package query

import (
	"strings"
)

type UserLikeQuery struct {
	userIDs  []uint64
	imageIDs []uint64
	cursor   *uint64
	pageSize *uint32
	orderBy  *string
}

func NewUserLikeQuery() *UserLikeQuery {
	return &UserLikeQuery{}
}

func (q *UserLikeQuery) UserID(
	userID uint64,
) *UserLikeQuery {
	q.userIDs = append(q.userIDs, userID)

	return q
}

func (q *UserLikeQuery) ImageID(
	imageID uint64,
) *UserLikeQuery {
	q.imageIDs = append(q.imageIDs, imageID)

	return q
}

func (q *UserLikeQuery) PageSize(
	pageSize *uint32,
) *UserLikeQuery {
	q.pageSize = pageSize

	return q
}

func (q *UserLikeQuery) Cursor(
	cursor *uint64,
) *UserLikeQuery {
	q.cursor = cursor

	return q
}

func (q *UserLikeQuery) OrderBy(
	orderBy *string,
) *UserLikeQuery {
	q.orderBy = orderBy

	return q
}

func (q *UserLikeQuery) Build() (wheres string, args []interface{}) {
	whereCols := make([]string, 0)

	if len(q.userIDs) != 0 {
		inCondition := "user_id IN (?"

		for i := 1; i < len(q.userIDs); i++ {
			inCondition = inCondition + ",?"
		}
		inCondition = inCondition + ")"
		whereCols = append(whereCols, inCondition)

		for _, id := range q.userIDs {
			args = append(args, id)
		}
	}

	if len(q.imageIDs) != 0 {
		inCondition := "image_id IN (?"

		for i := 1; i < len(q.imageIDs); i++ {
			inCondition = inCondition + ",?"
		}
		inCondition = inCondition + ")"
		whereCols = append(whereCols, inCondition)

		for _, id := range q.imageIDs {
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
