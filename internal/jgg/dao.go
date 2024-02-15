package jgg

import (
	"awesomeProject/db"
	"database/sql"
	"log"
)

const insertGe = `INSERT INTO 
    person_jgg (solar_date, lunar_date, leap_month, hour, solar_ge, lunar_ge, Element)
	VALUES (?, ?, ?, ?, ?, ?, ?)`

const selectGe = `SELECT 
    id, solar_date, lunar_date, leap_month, hour, solar_ge, lunar_ge, Element 
	FROM person_jgg`

const deleteGe = `DELETE FROM person_jgg where id = ?`

type DAO struct {
}

func (dao DAO) AddGe(ge *Ge) error {
	stmt, err := db.Inst.Prepare(insertGe)
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err = stmt.Close()
		if err != nil {
			log.Println(err)
		}
	}(stmt)

	_, err = stmt.Exec(ge.SolarDate, ge.LunarDate, ge.LeapMonth, ge.Hour, ge.Solar, ge.Lunar, ge.Element)
	if err != nil {
		return err
	}

	return nil
}

func (dao DAO) ListGe() ([]Ge, error) {
	query, err := db.Inst.Query(selectGe)
	if err != nil {
		return nil, err
	}
	defer func(query *sql.Rows) {
		err = query.Close()
		if err != nil {
			log.Println(err)
		}
	}(query)

	result := make([]Ge, 0)

	for query.Next() {
		ge := Ge{}
		err = query.Scan(&ge.Id, &ge.SolarDate, &ge.LunarDate, &ge.LeapMonth, &ge.Hour, &ge.Solar, &ge.Lunar, &ge.Element)
		if err != nil {
			return nil, err
		}
		result = append(result, ge)
	}

	return result, nil
}

func (dao DAO) DelGe(id int) (bool, error) {
	stmt, err := db.Inst.Prepare(deleteGe)
	if err != nil {
		return false, err
	}
	defer func(stmt *sql.Stmt) {
		err = stmt.Close()
		if err != nil {
			log.Println(err)
		}
	}(stmt)

	res, err := stmt.Exec(id)
	if err != nil {
		return false, err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	return affected > 0, nil
}
