package controllers

import (
	"fmt"
	"notable_health/cmd/main/models"
	"notable_health/pckg/db"
)

func InsertPhysician(physicians *models.Physicians) error {
	db := db.GetDB()

	query, err := db.Prepare(
		"INSERT INTO physicians " +
			"(physician_first_name, physician_last_name) " +
			"VALUES (?, ?)",
	)

	if err != nil {
		return fmt.Errorf("failed to create insert query")
	}

	_, err = query.Exec(physicians.FirstName, physicians.LastName)

	if err != nil {
		return fmt.Errorf("failed to insert to database")
	}

	return nil
}
