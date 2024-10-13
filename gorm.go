package pggorm

import (
	"errors"
	"gorm.io/gorm"
)

var (
	ConnectionErr     = errors.New("ConnectionErr")
	ErrRecordNotFound = errors.New("ErrRecordNotFound")
	ZeroRowsAffected  = errors.New("ZeroRowsAffected")
)

func Exist(err error, msg string) error {
	if err == nil {
		return errors.New(msg)
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return ConnectionErr
	}

	return nil
}

func NotExist(err error) error {
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}

	return ConnectionErr
}

func Delete(tx *gorm.DB) error {
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return ZeroRowsAffected
	}

	return nil
}

func Save(db *gorm.DB, T interface{}, isCreate bool) error {
	if isCreate {
		tx := db.Create(T)
		if tx.Error != nil {
			return tx.Error
		}
	} else {
		tx := db.Model(T).Updates(T)
		if tx.Error != nil {
			return tx.Error
		}

		if tx.RowsAffected == 0 {
			return ZeroRowsAffected
		}
	}

	return nil
}
