package database

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"sellease-ai/config"

	"time"

	_ "github.com/lib/pq"
	"github.com/pressly/goose"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func PrepareDatabase() (*gorm.DB, error) {
	config := config.GetConfig()

	dbURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		config.PostgresHost,
		config.PostgresPort,
		config.PostgresUsername,
		config.PostgresDB,
		config.PostgresPassword,
	)
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   config.PostgresSchema + ".",
			SingularTable: true,
		},
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		fmt.Println("error connecting database - %w", err)
		return nil, err
	}
	fmt.Println("db connection established!")

	err = autoMigrate(dbURL)
	if err != nil {
		log.Printf("error running db migration, skipping db migration - %v", err)
	}
	return db, nil
}

var (
	flags = flag.NewFlagSet("goose", flag.ExitOnError)
	dir   = flags.String("dir", "./database/migrations", "directory with migration files")
)

func autoMigrate(dbURL string) error {
	config := config.GetConfig()
	dbURL = fmt.Sprintf("%s search_path=%s", dbURL, config.PostgresSchema)

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return err
	}

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	if err := goose.Up(db, *dir); err != nil {
		return err
	}

	log.Println("db migration completed!")
	return nil
}
