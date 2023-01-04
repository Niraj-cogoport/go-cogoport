package migrations

import (
	"database/sql"
	_"gorm.io/driver/postgres"

)

func init() {
	migrator.AddMigration(&Migration{
		Version: "20230103201407",
		Up:      mig_20230103201407_init_schema_up,
		Down:    mig_20230103201407_init_schema_down,
	})
}

func mig_20230103201407_init_schema_up(tx *sql.Tx) error {
	_, err := tx.Exec("CREATE TABLE tickets ( name varchar(255) );")
	if err != nil {
		return err
	}
	return nil
}

func mig_20230103201407_init_schema_down(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE tickets;")
	if err != nil {
		return err
	}
	return nil
}