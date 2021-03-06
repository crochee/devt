package mysql

import (
	"context"

	"github.com/crochee/devt/internal/store"
	"github.com/crochee/devt/pkg/storage/mysql"
)

var (
	dbClient *mysql.DB
)

// Init init database
func Init(ctx context.Context) (err error) {
	dbClient, err = mysql.New(ctx)
	return
}

// DB 若想使用SELECT打印日志，请使用DB(ctx,database.WithLog())
func DB(ctx context.Context, opts ...mysql.Opt) *mysql.DB {
	return dbClient.With(ctx, opts...)
}

// GetMysqlFactory create mysql factory with context.Context
func GetMysqlFactory(ctx context.Context) store.Factory {
	return &dataStore{db: DB(ctx)}
}

type dataStore struct {
	db *mysql.DB
}

func (d *dataStore) Begin() store.Factory {
	d.db.Begin()
	return &dataStore{db: &mysql.DB{
		DB:    d.db.Begin(),
		Debug: d.db.Debug,
	}}
}

func (d *dataStore) Commit() {
	d.db.Commit()
}

func (d *dataStore) Rollback() {
	d.db.Rollback()
}

func (d *dataStore) Auth() store.AuthorControlStore {
	return newAuthorControl(d.db)
}

func (d *dataStore) Flow() store.ChangeFlowStore {
	return newResourceChangeFlow(d.db)
}

func (d *dataStore) Pkg() store.ResourcePkgStore {
	return newResourcePkg(d.db)
}
