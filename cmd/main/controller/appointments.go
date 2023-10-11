package controller

import (
	"database/sql"
	"fmt"
	"notable_health/cmd/main/models"
	"notable_health/pckg/db"
)

func GetPhysicianList() ([]models.PhysicianListData, error) {
	db := db.GetDB()

	getQuery := "SELECT * FROM physicians"

	rows, err := db.Query(getQuery)

	if err != nil {
		return nil, fmt.Errorf("error querying database")
	}

	var physicianList []models.PhysicianListData

	for rows.Next() {
		var physicianData models.PhysicianListData

		if err := rows.Scan(
			&physicianData.IdPhysician,
			&physicianData.FirstName,
			&physicianData.LastName,
		); err != nil {
			return nil, fmt.Errorf("error scanning row data")
		}

		physicianList = append(physicianList, physicianData)
	}

	return physicianList, nil
}

func CheckValidPhysician(reqBody *models.InsertPhysicianData) error {
	db := db.GetDB()

	getQuery := "SELECT lastName FROM physicians " +
		"WHERE firstName = ? AND lastName = ?"

	var physician models.InsertPhysicianData

	err := db.QueryRow(getQuery, reqBody.FirstName, reqBody.LastName).Scan(&physician.LastName)

	fmt.Println("this is err!! ", err)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}

		return err
	}

	return fmt.Errorf("physician already exists")
}

func InsertPhysicianData(reqBody *models.InsertPhysicianData) error {
	db := db.GetDB()

	insertQuery := "INSERT INTO physicians (firstName, lastName) " +
		"VALUES (?, ?)"

	_, err := db.Exec(insertQuery, reqBody.FirstName, reqBody.LastName)

	if err != nil {
		return fmt.Errorf("error inserting data to database")
	}

	return nil
}
