package controller

import (
	"database/sql"
	"fmt"
	"notable_health/cmd/main/models"
	"notable_health/pckg/db"
	"strings"
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

func InsertAppointmentData(reqBody *models.AddAppointmentData, id int) error {
	db := db.GetDB()

	formattedDate := dateSplit(reqBody.Date_Column)

	checkQuery := "SELECT patientLastName FROM appointments " +
		"WHERE date_column = ? AND time_column = ? AND IdPhysician = ?"

	insertQuery := "INSERT INTO appointments " +
		"(patientFirstName, patientLastName, date_column, time_column, kind, IdPhysician) " +
		"VALUE (?, ?, ?, ?, ?, ?)"

	rows, err := db.Query(
		checkQuery,
		formattedDate,
		reqBody.Time_Column,
		id,
	)

	if err != nil {
		return fmt.Errorf("error querying for appointment data")
	}

	var appointments []models.CheckAppointmentData

	for rows.Next() {
		var appointmentData models.CheckAppointmentData

		if err := rows.Scan(
			&appointmentData.PatientLastName,
		); err != nil {
			return fmt.Errorf("error scanning rows")
		}

		appointments = append(appointments, appointmentData)
	}

	if len(appointments) > 2 {
		return fmt.Errorf("too many appointments at this time")
	}

	_, err = db.Exec(
		insertQuery,
		reqBody.PatientFirstName,
		reqBody.PatientLastName,
		formattedDate,
		reqBody.Time_Column,
		reqBody.Kind,
		id,
	)

	if err != nil {
		return fmt.Errorf("error inserting data into database")
	}

	return nil

}

func DeleteAppointment(IdApp int, IdPhy int) error {
	db := db.GetDB()

	deleteQuery := "DELETE FROM appointments WHERE IdAppointment = ? AND IdPhysician = ?"

	_, err := db.Exec(
		deleteQuery,
		IdApp,
		IdPhy,
	)

	if err != nil {
		return fmt.Errorf("error executing delete query")
	}

	return nil
}

func dateSplit(date string) string {
	dateParts := strings.Split(date, "/")

	return dateParts[2] + "-" + dateParts[1] + "-" + dateParts[0]
}
