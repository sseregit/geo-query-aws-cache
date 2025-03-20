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
	var res types.User

	var image interface{}
	var hobby interface{}

	if err := d.db.QueryRow(GetUserByName, userName).Scan(&res.UserName, &image, &res.Description, &hobby, &res.Latitude, &res.Hardness); err != nil {
		return nil, err
	} else if err = unMarshalToField(
		[]interface{}{image, hobby},
		&res.Image, &res.Hobby,
	); err != nil {
		return nil, err
	} else {
		return &res, nil
	}
}

func (d *DB) AroundUser(userName string, latitude, hardness float64, searchRange, limit int64) ([]*types.User, error) {
	if rows, err := d.db.Query(GetAroundUsers, userName, hardness, latitude, searchRange, hardness, latitude, limit); err != nil {
		return nil, err
	} else {
		defer rows.Close()

		var result []*types.User

		for rows.Next() {
			var res types.User

			var image interface{}
			var hobby interface{}

			if err := rows.Scan(&res.UserName, &image, &res.Description, &hobby, &res.Latitude, &res.Hardness); err != nil {
				return nil, err
			} else if err = unMarshalToField(
				[]interface{}{image, hobby},
				&res.Image, &res.Hobby,
			); err != nil {
				return nil, err
			} else {
				result = append(result, &res)
			}
		}
		return result, nil
	}
}
