package query

import (
	"fmt"
	"strings"
)

type UserQuery struct {
	ids     []uint64
	emails  []string
	keyword *string
}

func NewUserQuery() *UserQuery {
	return &UserQuery{}
}

func (q *UserQuery) ID(
	ID uint64,
) *UserQuery {
	q.ids = append(q.ids, ID)

	return q
}

func (q *UserQuery) IDs(
	IDs []uint64,
) *UserQuery {
	q.ids = append(q.ids, IDs...)

	return q
}

func (q *UserQuery) Email(
	email string,
) *UserQuery {
	q.emails = append(q.emails, email)

	return q
}

func (q *UserQuery) Emails(
	email []string,
) *UserQuery {
	q.emails = append(q.emails, email...)

	return q
}

func (q *UserQuery) Keyword(
	keyword *string,
) *UserQuery {
	q.keyword = keyword

	return q
}

func (q *UserQuery) Build() (wheres string, args []interface{}) {
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

	if len(q.emails) != 0 {
		inCondition := "email_address IN (?"

		for i := 1; i < len(q.emails); i++ {
			inCondition = inCondition + ",?"
		}
		inCondition = inCondition + ")"
		whereCols = append(whereCols, inCondition)

		for _, userName := range q.emails {
			args = append(args, userName)
		}
	}

	if q.keyword != nil {
		whereCols = append(whereCols,
			fmt.Sprintf("username LIKE '%s%%'", *q.keyword),
		)
	}

	wheres = strings.Join(whereCols, " AND ")

	return wheres, args
}
