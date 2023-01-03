package migrations

import (
	"database/sql"
)

func init() {
	migrator.AddMigration(&Migration{
		Version: "20230103183537",
		Up:      mig_20230103183537_init_schema_up,
		Down:    mig_20230103183537_init_schema_down,
	})
}

func mig_20230103183537_init_schema_up(tx *sql.Tx) error {
	_, err := tx.Exec("CREATE TABLE groups (id uuid PRIMARY KEY,name varchar,tags varchar,status varchar,created_at datetime,updated_at datetime;")
	if err != nil {
		return err
	}
	return nil
}

func mig_20230103183537_init_schema_down(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE groups")
	if err != nil {
		return err
	}
	return nil
}