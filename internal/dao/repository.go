package dao

import (
	"github.com/orderoutofchaos/oj-server/pkg/sq"
)

var repo Repository

type Repository interface {
	Writer() *sq.DB
	Reader() *sq.DB
}

type repositoryImpl struct {
	writer *sq.DB
	reader *sq.DB
}

func (r repositoryImpl) Writer() *sq.DB {
	return r.writer
}

func (r repositoryImpl) Reader() *sq.DB {
	return r.reader
}

func InitRepo(reader *sq.DB, writer *sq.DB) {
	repo = repositoryImpl{
		writer: writer,
		reader: reader,
	}
}

func GetRepo() Repository {
	return repo
}
