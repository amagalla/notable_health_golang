package controllers

import (
	"fmt"
	"notable_health/cmd/main/models"
	"notable_health/pckg/db"
)

func InsertPhysician(physicians *models.PostPhysiciansData) error {
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

func GetPhysicianList() ([]models.GetPhysicanData, error) {
	db := db.GetDB()

	rows, err := db.Query("SELECT * FROM physicians")

	if err != nil {
		return nil, fmt.Errorf("error querying for list of physicians")
	}

	var physicianList []models.GetPhysicanData

	for rows.Next() {
		var physicanData models.GetPhysicanData

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

func PostAppointmentData(appointment *models.PostAppointmentData) error {
	db := db.GetDB()

	query, err := db.Prepare(
		"INSERT INTO appointment " +
			"(patient_first_name, patient_last_name, scheduled_date, " +
			"scheduled_time, kind, physician_id) " +
			"VALUES (?, ?, ?, ?, ?, ?)",
	)

	if err != nil {
		return fmt.Errorf("error preparing query")
	}

	_, err = query.Exec(
		appointment.PatientFirstName,
		appointment.PatientLastName,
		appointment.ScheduledDate,
		appointment.ScheduledTime,
		appointment.Kind,
		appointment.PhysicianID,
	)

	if err != nil {
		return fmt.Errorf("error executing query")
	}

	return nil
}
