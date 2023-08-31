package gorm

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func MakeSureDBTableExist(db *gorm.DB, table interface{}) error {
	if db == nil {
		return errors.New("db required")
	}
	err := db.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&table)
	if err != nil {
		return err
	}
	return nil
}

func AddOneRecord(db *gorm.DB, dto interface{}) error {
	if db == nil {
		return errors.New("db required")
	}

	return db.Create(dto).Error
}

func DelOneRecord(db *gorm.DB, dto interface{}, whereQuery interface{}, whereArgs ...interface{}) error {
	if db == nil {
		return errors.New("db required")
	}

	tx := db.Delete(dto, whereQuery, whereArgs)
	if tx.Error != nil {
		return tx.Error
	}

	rowsAffected := tx.RowsAffected
	if rowsAffected != 0 {
		return nil
	}
	return errors.New("empty record")
}

func UpdateOneRecord(db *gorm.DB, dto interface{}) error {
	if db == nil {
		return errors.New("db required")
	}

	// if exists, update; else insert
	return db.Save(dto).Error
}

func QueryOneRecord(db *gorm.DB, dto interface{}, whereQuery interface{}, whereArgs ...interface{}) error {
	if db == nil {
		return errors.New("db required")
	}

	tx := db.Where(whereQuery, whereArgs...).First(dto)
	if tx.Error != nil {
		return tx.Error
	}

	rowsAffected := tx.RowsAffected
	if rowsAffected != 0 {
		return nil
	}
	return errors.New("empty record")
}

func UpdateOneColumn(db *gorm.DB, dbDTO interface{}, columnName string, columnValue interface{},
	whereQuery interface{}, whereArgs ...interface{}) error {
	if db == nil {
		return errors.New("db required")
	}

	tx := db.Model(dbDTO).Where(whereQuery, whereArgs...).Update(columnName, columnValue)
	if tx.Error != nil {
		return tx.Error
	}

	rowsAffected := tx.RowsAffected
	if rowsAffected != 0 {
		return nil
	}
	return fmt.Errorf("empty record or same %v", columnName)
}
