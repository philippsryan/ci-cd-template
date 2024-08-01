package db

import (
	"database/sql"
	"errors"
	"io/fs"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Migration struct {
	id   int
	sql  string
	name string
}

func (m Migration) Apply(db *sql.DB) error {
	_, err := db.Exec(m.sql)
	return err
}

type ByMigrationFileNumber []fs.DirEntry

func (mf ByMigrationFileNumber) Len() int      { return len(mf) }
func (nf ByMigrationFileNumber) Swap(i, j int) { nf[i], nf[j] = nf[j], nf[i] }
func (mf ByMigrationFileNumber) Less(i, j int) bool {

	name_components := strings.Split(mf[i].Name(), ".")
	i_id, _ := strconv.Atoi(name_components[1])

	name_components = strings.Split(mf[j].Name(), ".")
	j_id, _ := strconv.Atoi(name_components[1])

	return i_id < j_id
}

func loadMigrationFiles() ([]Migration, error) {
	migration_files, err := os.ReadDir("./migrations")

	if err != nil {
		println("Failed finding migration folder: ", err.Error())
		return nil, err
	}

	sort.Sort(ByMigrationFileNumber(migration_files))
	migrations := make([]Migration, len(migration_files))

	for _, file := range migration_files {
		var migration Migration

		name_components := strings.Split(file.Name(), ".")

		migration.name = name_components[0]
		migration.id, err = strconv.Atoi(name_components[1])

		if err != nil {
			println("Failed to parse migration file name for information")
			return nil, err
		}

		sql_query, err := os.ReadFile("./migrations/" + file.Name())

		if err != nil {
			println("Failed to read migration file")
			return nil, err
		}

		migration.sql = string(sql_query)

		migrations = append(migrations, migration)
	}

	return migrations, nil
}

func doesMigrationTableExists(db *sql.DB) bool {
	top_row, sql_error := db.Query("SELECT * FROM Migrations LIMIT 1;")

	if sql_error != nil {
		return false
	} else {
		return top_row.Next()
	}
}

func createMigrationTable(db *sql.DB) error {

	migration_table_qry := "CREATE TABLE Migrations (Current_Migration int, Date_Performed DATE);"

	_, err := db.Exec(migration_table_qry)

	if err != nil {
		return err
	}

	insert_result, err := db.Exec("INSERT INTO Migration (Current_Migration, Date_Performed) VALUES (NULL, NULL);")

	if err != nil {
		return err
	}

	_, err = insert_result.RowsAffected()

	if err != nil {
		return err
	}

	return nil
}

func RunMigrations(db *sql.DB) (int, error) {
	/*
		1. Load in all migration sql files from the migration folder
		2. Sort migrations based on there file name
		3. Check If the migration table exists
			a. If it doesn't create it
		4. Check the last migration id that was ran
		5. Run every migration in order that is after it
		6. Save the last migration id to the migration table
	*/

	migrations, err := loadMigrationFiles()

	if err != nil {
		return 0, err
	}

	sort.Slice(migrations, func(i, j int) bool {
		return migrations[i].id < migrations[j].id
	})

	have_migration_table := doesMigrationTableExists(db)

	if !have_migration_table {
		// create the migration table
		createMigrationTable(db)
	}

	// get the current migration

	var current_migration int

	row := db.QueryRow("SELECT * FROM Migrations LIMIT 1;")

	if row.Err() != nil {
		return 0, row.Err()
	}

	row.Scan(&current_migration)

	migrations_completed := 0
	last_completed_migration := current_migration

	for _, migration := range migrations {
		if migration.id < current_migration {
			continue
		}

		err := migration.Apply(db)

		if err != nil && last_completed_migration != current_migration {
			_, update_err := db.Exec("UPDATE Migrations SET Current_Migration = ?, Date_Performed = CURDATE();", migration.id)

			if update_err != nil {
				return 0, errors.Join(err, update_err)
			}

			return migrations_completed, err
		}

		migrations_completed += 1
		last_completed_migration = migration.id
	}

	db.Exec("UPDATE Migrations SET Current_Migration = ?, SET Date_Performed = CURRENT_DATE();", last_completed_migration)

	return migrations_completed, nil
}
