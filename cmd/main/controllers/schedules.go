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

func GetPhysicianList() ([]models.PhysicanData, error) {
	db := db.GetDB()

	rows, err := db.Query("SELECT * FROM physicians")

	if err != nil {
		return nil, fmt.Errorf("error querying for list of physicians")
	}

	var physicianList []models.PhysicanData

	for rows.Next() {
		var physicanData models.PhysicanData

		if err := rows.Scan(
			&physicanData.PhysicanID,
			&physicanData.FirstName,
			&physicanData.LastName,
		); err != nil {
			return nil, fmt.Errorf("error scanning rows")
		}

		physicianList = append(physicianList, physicanData)
	}

	return physicianList, nil
}
