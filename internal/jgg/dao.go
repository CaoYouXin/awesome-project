package jgg

import (
	"awesomeProject/db"
	"database/sql"
	"log"
)

type DAO struct {
}

func (dao DAO) AddGe(ge *Ge) error {
	stmt, err := db.Inst.Prepare("INSERT INTO person_jgg (solar_date, lunar_date, leap_month, hour, solar_ge, lunar_ge, Element) VALUES (?, ?, ?, ?, ?, ?, ?)")
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
	query, err := db.Inst.Query("SELECT id, solar_date, lunar_date, leap_month, hour, solar_ge, lunar_ge, Element FROM person_jgg")
	if err != nil {
		return nil, err
	}
	defer func(query *sql.Rows) {
		err = query.Close()
		if err != nil {
			log.Println(err)
		}
	}(query)

	var res []Ge

	for query.Next() {
		ge := Ge{}
		err = query.Scan(&ge.Id, &ge.SolarDate, &ge.LunarDate, &ge.LeapMonth, &ge.Hour, &ge.Solar, &ge.Lunar, &ge.Element)
		if err != nil {
			return nil, err
		}
		res = append(res, ge)
	}

	return res, nil
}

func (dao DAO) DelGe(id int) (bool, error) {
	stmt, err := db.Inst.Prepare("DELETE FROM person_jgg where id = ?")
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
