package repo

import (
	"log"

	"github.com/LorezV/gophkeeper/internal/client/usecase/repo/models"
	"github.com/fatih/color"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type GophKeeperRepo struct {
	db *gorm.DB
}

func New(dbFileName string) *GophKeeperRepo {
	db, err := gorm.Open(sqlite.Open(dbFileName), &gorm.Config{})
	if err != nil {
		color.Red("Load error %s", err.Error())
	}

	return &GophKeeperRepo{
		db: db,
	}
}

func (r *GophKeeperRepo) MigrateDB() {
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

	var err error

	for _, table := range tables {
		if err = r.db.Migrator().DropTable(table); err != nil {
			log.Fatalf("Init error %s", err.Error())
		}

		if err = r.db.Migrator().CreateTable(table); err != nil {
			log.Fatalf("Init error %s", err.Error())
		}
	}

	color.Green("Initialization status: success")
	color.Green("You can use gophkeer")
}
