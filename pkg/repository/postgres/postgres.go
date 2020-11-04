package postgres

import (
	"github.com/jessalva/go-railwire/pkg/repository"
	"github.com/jessalva/go-railwire/pkg/repository/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type postgresDB struct {
	db *gorm.DB
}



func NewPostgres() repository.PlanRepository {

	newPostgresDB := postgresDB{}

	db, err := gorm.Open( postgres.Open("host=postgres port=5432 user=postgres password=railsa dbname=railwire host=postgres sslmode=disable"), &gorm.Config{} )
	newPostgresDB.db = db

	if err != nil {
		log.Fatal( "Postgres not located" )
	}

	_ = newPostgresDB.db.AutoMigrate(&models.FUPPlan{})

	return newPostgresDB

}