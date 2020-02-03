package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/jgolang/mysqltools"
)

// User ...
type User struct {
	ID             int64
	FirstName      string
	FirstLastName  string
	SecondLastName string
	Email          string
	Phone          string
	Address        string
	Sons           []Son
}

// Son ...
type Son struct {
	ID             int64
	FirstName      string
	FirstLastName  string
	SecondLastName string
	Avatar         int64
	Section        string
	Grade          string
}

func validateUser(email string) error {

	query := fmt.Sprintf("SELECT id FROM users WHERE email = @email")
	query, err := mysqltools.GetQueryString(
		query,
		sql.Named("email", email),
	)
	if err != nil {
		return err
	}

	row := db.QueryRow(query)

	var id int64

	err = row.Scan(&id)
	if err != nil {
		return err
	}

	if id == 0 {
		return errors.New("No se ha encontrado el usuario")
	}

	return nil
}

func getUserInfo(email string) (user User, err error) {

	query := fmt.Sprintf("SELECT id, first_name, first_last_name, second_last_name, email, phone, address FROM mas_person WHERE email = @email")
	query, err = mysqltools.GetQueryString(
		query,
		sql.Named("email", email),
	)
	if err != nil {
		return
	}

	row := db.QueryRow(query)

	err = row.Scan(
		&user.ID,
		&user.FirstName,
		&user.FirstLastName,
		&user.SecondLastName,
		&user.Email,
		&user.Phone,
		&user.Address,
	)

	if err != nil {
		return
	}

	return
}

func getUserSons(userID int64) (sons []Son, err error) {
	var query2 string
	query := `
		SELECT 
			O.classmate, 
			P.first_name, 
			P.first_last_name, 
			P.second_last_name, 
			P.avatar,
			(
				SELECT  
					ms.name
				FROM 
					assignation a
				JOIN 
					mas_period mp 
						ON a.period_id = mp.id 
						AND mp.current = 1 
						AND mp.deleted_at IS NULL
				JOIN 
					section s 
						ON s.id = a.section_id 
						AND s.deleted_at IS NULL
				JOIN 
					mas_section ms 
						ON ms.id = s.mas_section_id 
						AND ms.deleted_at IS NULL
				WHERE 
					a.person_id = P.id  
			) AS section, 
			(
				SELECT  
					mg.description
				FROM 
					assignation a
				JOIN 
					mas_period mp 
						ON a.period_id = mp.id 
						AND mp.current = 1 
						AND mp.deleted_at IS NULL
				JOIN 
					section s 
						ON s.id = a.section_id 
						AND s.deleted_at IS NULL
				JOIN 
					mas_section ms 
						ON ms.id = s.mas_section_id 
						AND ms.deleted_at IS NULL
				JOIN 
					mas_grade mg 
						ON mg.id = ms.grade_id
				WHERE 
					a.person_id = P.id  
			) AS grade
		FROM 
			owner_classmate AS O 
		INNER JOIN 
			mas_person AS P 
				ON O.classmate = P.id 
		WHERE 
			O.owner = @userID 
			AND O.deleted_at IS NULL
	`
	query2, err = mysqltools.GetQueryString(
		query,
		sql.Named("userID", userID),
	)
	if err != nil {
		return
	}

	log.Println(query2)

	rows, errR := db.Query(query2)
	if errR != nil {
		return sons, errR
	}

	for rows.Next() {
		var son Son
		err = rows.Scan(
			&son.ID,
			&son.FirstName,
			&son.FirstLastName,
			&son.SecondLastName,
			&son.Avatar,
			&son.Section,
			&son.Grade,
		)
		if err != nil {
			return
		}

		sons = append(sons, son)
	}

	return
}
