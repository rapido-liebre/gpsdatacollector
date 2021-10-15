package database

import (
	"context"
	"fmt"
	"math"

	"github.com/go-kit/kit/log"
	"github.com/rapido-liebre/gpsDataCollector/internal"
	"github.com/rapido-liebre/gpsDataCollector/pkg/receiver"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type repo struct {
	db     *gorm.DB
	logger log.Logger
}

func NewRepo(db *gorm.DB, logger log.Logger) receiver.Repository {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "sql"),
	}
}

func (r repo) AddCoordinates(ctx context.Context, data *internal.GpsData) error {
	var result internal.GpsData
	r.db.Raw(
		"insert into coordinates (DEVICE_ID,POINT) values ( ?, point(?, ?) )",
		data.DeviceId,
		math.Round(float64(data.Point.Latitude*100))/100,
		math.Round(float64(data.Point.Longitude*100))/100,
	).Scan(&result)

	return r.db.Error
}

func (r repo) deleteLastRecFromCoordinates() error {
	r.db.Raw("delete from coordinates order by id desc limit 1")
	return r.db.Error
}

func Init(host, port, dbname, user, pass string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, pass, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
	//defer db.Close()
}
