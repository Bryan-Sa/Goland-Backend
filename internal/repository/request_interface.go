package repository

import (
	"database/sql"
	"reflect"

	"github.com/Bryan-Sa/Goland-Backend/db"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// true find all
func FindAll(table string, entities interface{}) ([]interface{}, error) {
	// récupérer les lignes
	rows, err := db.DB.Query("SELECT * FROM " + table)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// récupérer les colonnes par rapport aux lignes
	columns, err_column := rows.Columns()
	if err_column != nil {
		return nil, err_column
	}
	// Créer un slice de pointeurs pour les champs de l'entité (Voir si je peux mettre tout ce qu'il ya en bas dans une autre fonction)
	entityType := reflect.TypeOf(entities).Elem()
	resultStruct := make([]interface{}, len(columns))
	for i := range columns {
		field := entityType.Field(i)
		resultStruct[i] = reflect.New(field.Type).Interface()
	}
	var result []interface{}
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		if err := rows.Scan(resultStruct...); err != nil {
			return result, err
		}
		entity := reflect.New(entityType).Elem()
		for i := range columns {
			field := entity.Field(i)
			field.Set(reflect.ValueOf(resultStruct[i]).Elem())
		}
		result = append(result, entity.Interface())
	}
	if err = rows.Err(); err != nil {
		return result, err
	}
	return result, nil
}

func FindAllByKey(table string, entities interface{}) []interface{} {

}

// Ajoutez d'autres fonctions de repository selon vos besoins.
