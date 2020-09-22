package postgres

import(
	"fmt"
	"context"
	"github.com/go-pg/pg/v10"
)

type DBLogger struct {

}

func (d DBLogger) BeforeQuery(ctx context.Context, q *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}

func (d DBLogger) AfterQuery(ctx context.Context, q *pg.QueryEvent) (error) {
	s, _ := q.FormattedQuery()
	fmt.Println(string(s))
	return nil
}

func New(opts *pg.Options) *pg.DB {
	return  pg.Connect(opts)
}