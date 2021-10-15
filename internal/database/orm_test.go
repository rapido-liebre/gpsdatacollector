package database

import (
	"context"
	"os"
	"testing"

	//"github.com/rapido-liebre/gpsDataCollector/config"
	"github.com/go-kit/kit/log"
	"github.com/rapido-liebre/gpsDataCollector/internal"
	"gorm.io/gorm"
)

func TestInit(t *testing.T) {
	type args struct {
		host   string
		port   string
		dbname string
		user   string
		pass   string
	}
	tests := []struct {
		name    string
		args    args
		want    *gorm.DB
		wantErr bool
	}{
		{"Valid connection", args{"localhost", "5436", "gps", "admin", "password"}, nil, false},
		{"Invalid port", args{"localhost", "5432", "gps", "admin", "password"}, nil, true},
		{"Invalid db name", args{"localhost", "5436", "foo", "admin", "password"}, nil, true},
		{"Invalid credentials", args{"localhost", "5436", "gps", "foo", "bar"}, nil, true},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				_, err := Init(tt.args.host, tt.args.port, tt.args.dbname, tt.args.user, tt.args.pass)
				if (err != nil) != tt.wantErr {
					t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				//if !reflect.DeepEqual(got, tt.want) {
				//	t.Errorf("Init() got = %v, want %v", got, tt.want)
				//}
			},
		)
	}
}

func Test_repo_AddCoordinates(t *testing.T) {
	//test setup
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	db, err := Init("localhost", "5436", "gps", "admin", "password")
	if err != nil {
		t.Errorf("Init() error = %v", err)
		return
	}

	type fields struct {
		db     *gorm.DB
		logger log.Logger
	}
	type args struct {
		ctx  context.Context
		data *internal.GpsData
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"Insert coordinates successful",
			fields{db, logger}, args{
			ctx:  context.Background(),
			data: &internal.GpsData{
				Id:       "",
				DeviceId: "foo",
				Point:    internal.Coordinates{
					Latitude:  77.345,
					Longitude: 99.234,
				},
			},
		}, false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				r := repo{
					db:     tt.fields.db,
					logger: tt.fields.logger,
				}
				if err := r.AddCoordinates(tt.args.ctx, tt.args.data); (err != nil) != tt.wantErr {
					t.Errorf("AddCoordinates() error = %v, wantErr %v", err, tt.wantErr)
				}
				//cleanup
				if err := r.deleteLastRecFromCoordinates(); err != nil {
					t.Errorf("deleteLastRecFromCoordinates() error = %v", err)
				}
			},
		)
	}
}
