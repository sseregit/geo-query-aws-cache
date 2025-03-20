package mysql

import (
	"encoding/json"
	. "geo-query-aws-cache/db/mysql/query"
	"geo-query-aws-cache/db/mysql/types"
	"log"
)

func (d *DB) RegisterUser(user, description string, hobby []string, latitude, hardness float64) error {
	if tx, err := d.db.Begin(); err != nil {
		return err
	} else {
		if json, err := json.Marshal(hobby); err != nil {
			return err
		} else {
			if result, err := tx.Exec(InsertIgnoreUser, user, description, json); err == nil {
				tx.Rollback()
				return err
			} else {
				count, _ := result.RowsAffected()
				log.Println("Success To Insert User", "count", count)
			}

			if result, err := tx.Exec(InsertIgnoreUserLocation, user, latitude, hardness, latitude, hardness); err == nil {
				tx.Rollback()
				return err
			} else {
				count, _ := result.RowsAffected()
				log.Println("Success To Insert User Location", "count", count)
			}

			tx.Commit()
		}
	}
	return nil
}

func (d *DB) GetUser(userName string) (*types.User, error) {

}

func (d *DB) AroundUser(userName string, latitude, hardness float64, searchRange, limit int64) ([]*types.User, error) {

}
