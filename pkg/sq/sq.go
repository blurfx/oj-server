package sq

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
)

const fieldTag = "db"

type field struct {
	Name    string
	Type    reflect.Type
	Indices []int
}

type Rows struct {
	*sql.Rows
	fields map[string]*field
}

func (r *Rows) findFields(t reflect.Type) error {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	columnCheck := make(map[string]bool)

	fieldCount := t.NumField()
	for i := 0; i < fieldCount; i++ {
		f := t.Field(i)
		tag := f.Tag.Get(fieldTag)

		if tag == "" {
			continue
		}
		columnCheck[tag] = true

		fieldIndex := f.Index[0]
		fieldType := f.Type

		if fieldType.Kind() == reflect.Ptr {
			fieldType = fieldType.Elem()
		}

		r.fields[tag] = &field{
			Name:    tag,
			Type:    fieldType,
			Indices: []int{fieldIndex},
		}
	}

	columns, err := r.Columns()
	if err != nil {
		return err
	}
	for _, column := range columns {
		if !columnCheck[column] {
			return errors.New(fmt.Sprintf("No destination field found for column '%s'", column))
		}
	}
	return nil
}

func (r *Rows) ScanStruct(dest interface{}) error {
	v := reflect.ValueOf(dest)
	if v.Kind() != reflect.Ptr {
		return errors.New("Destination must be a pointer")
	}

	elem := v.Elem()
	if elem.Kind() != reflect.Struct {
		return errors.New(fmt.Sprintf("Element kind '%s' is not supported", elem.Kind()))
	}

	destType := reflect.TypeOf(dest)
	err := r.findFields(destType)
	if err != nil {
		return err
	}

	columns, err := r.Columns()
	if err != nil {
		return err
	}

	scanFieldPtrs := make([]interface{}, len(columns))
	scanFields := make([]*field, len(columns))
	for i, column := range columns {
		field := r.fields[column]
		scanFieldPtrs[i] = reflect.New(reflect.PtrTo(field.Type)).Interface()
		scanFields[i] = field
	}

	err = r.Scan(scanFieldPtrs...)
	if err != nil {
		return err
	}

	for i, field := range scanFieldPtrs {
		value := reflect.ValueOf(field).Elem().Elem()
		destField := elem
		for j := range scanFields[i].Indices {
			if destField.Kind() == reflect.Ptr {
				if destField.IsNil() {
					if !value.IsValid() {
						break
					}

					newValue := reflect.New(destField.Type().Elem())
					destField.Set(newValue)
				}
				destField = destField.Elem()
			}
			destField = destField.Field(scanFields[i].Indices[j])
		}

		if !value.IsValid() {
			destField.Set(reflect.Zero(destField.Type()))
		} else if destField.Kind() == reflect.Ptr {
			newValue := reflect.New(destField.Type().Elem())
			newValue.Elem().Set(value)
			destField.Set(newValue)
		} else if destField.Kind() == reflect.Struct {
			newValue := reflect.New(destField.Type())
			newValue.Elem().Set(value)
			destField.Set(value)
		} else {
			destField.Set(value)
		}
	}
	return nil
}

type Row struct {
	rows *Rows
	err  error
}

func (r *Row) Scan(dest ...any) error {
	defer r.rows.Close()
	if !r.rows.Next() {
		if r.rows.Err() != nil {
			return r.rows.Err()
		}
		return sql.ErrNoRows
	}
	return r.rows.Scan(dest...)
}

func (r *Row) ScanStruct(dest interface{}) error {
	defer r.rows.Close()
	if !r.rows.Next() {
		if r.rows.Err() != nil {
			return r.rows.Err()
		}
		return sql.ErrNoRows
	}
	return r.rows.ScanStruct(dest)
}

type DB struct {
	*sql.DB
	instance *sql.DB
}

func NewDb(db *sql.DB) *DB {
	return &DB{
		DB:       db,
		instance: db,
	}
}

func (db *DB) Query(query string, args ...any) (*Rows, error) {
	rows, err := db.instance.Query(query, args...)
	return &Rows{
		Rows:   rows,
		fields: make(map[string]*field),
	}, err
}

func (db *DB) QueryRow(query string, args ...any) *Row {
	rows, err := db.Query(query, args...)
	return &Row{
		rows: rows,
		err:  err,
	}
}
