package query

import (
	"fmt"
	"strings"
)

type UserQuery struct {
	ids       []uint64
	usernames []string
	emails    []string
	keyword   *string
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

func (q *UserQuery) Username(
	username string,
) *UserQuery {
	q.usernames = append(q.usernames, username)

	return q
}

func (q *UserQuery) Usernames(
	usernames []string,
) *UserQuery {
	q.usernames = append(q.usernames, usernames...)

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

		for _, email := range q.emails {
			args = append(args, email)
		}
	}

	if len(q.usernames) != 0 {
		inCondition := "username in (?"

		for i := 1; i < len(q.usernames); i++ {
			inCondition = inCondition + ",?"
		}
		inCondition = inCondition + ")"
		whereCols = append(whereCols, inCondition)

		for _, username := range q.usernames {
			args = append(args, username)
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
