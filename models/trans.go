package models

import (
	"gorm.io/gorm"
)

type Trans []func(tx *gorm.DB) error

func (d Trans) Exec() error {
	if len(d) == 0 {
		return nil
	}
	tx := db.Begin()
	for _, tran := range d {
		err := tran(tx)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}

func GetTx() *gorm.DB {
	tx := db.Begin()
	return tx
}
