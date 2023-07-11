package repo

import (
	"log"
	"time"

	"github.com/LorezV/gophkeeper/internal/server/usecase/repo/models"
	"github.com/LorezV/gophkeeper/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// GophKeeper - Repo.
type GophKeeperRepo struct {
	db *gorm.DB
	l  *logger.Logger
}

// New -.
func New(dsn string, l *logger.Logger) *GophKeeperRepo {
	attempts := 20
	for attempts > 0 {
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			return &GophKeeperRepo{
				db: db,
				l:  l,
			}
		}
		log.Printf("Database: %s is not available, attempts left: %d", dsn, attempts)
		time.Sleep(time.Second)
		attempts--
	}
	log.Fatalln("GophKeeperRepo - New - could not connect")

	return nil
}

func (r *GophKeeperRepo) Migrate() {
	tables := []interface{}{
		&models.User{},
		&models.Card{},
		&models.MetaCard{},
		&models.Login{},
		&models.MetaLogin{},
		&models.Note{},
		&models.MetaNote{},
		&models.Binary{},
		&models.MetaBinary{},
	}

	if err := r.db.AutoMigrate(tables...); err != nil {
		r.l.Fatal("GophKeeperRepo - Migrate - %v", err)
	}

	r.l.Debug("GophKeeperRepo - Migrate - success")
}

func (r *GophKeeperRepo) DBHealthCheck() error {
	sqlDB, err := r.db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Ping()
}

func (r *GophKeeperRepo) ShutDown() {
	db, err := r.db.DB()
	if err != nil {
		r.l.Error(err)
	}

	db.Close()
	r.l.Debug("db connection closed")
}
