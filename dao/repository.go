package dao

import (
    "database/sql"
)

var repo Repository

type Repository interface {
    Writer() *sql.DB
    Reader() *sql.DB
}

type repositoryImpl struct {
    writer *sql.DB
    reader *sql.DB
}

func (r repositoryImpl) Writer() *sql.DB {
    return r.writer
}

func (r repositoryImpl) Reader() *sql.DB {
    return r.reader
}

func InitRepo(reader *sql.DB, writer *sql.DB) {
    repo = repositoryImpl{
        writer: writer,
        reader: reader,
    }
}

func GetRepo() Repository {
    return repo
}
