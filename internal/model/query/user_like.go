package query

import (
	"strings"
)

type UserLikeQuery struct {
	userIDs  []uint64
	imageIDs []uint64
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

	wheres = strings.Join(whereCols, " AND ")

	return wheres, args
}
