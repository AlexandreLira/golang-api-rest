package auto

import (
	"log"

	"github.com/AlexandreLira/goApi/api/database"
	"github.com/AlexandreLira/goApi/api/models"
)

func Load() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		log.Fatal(err)
	}

	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatal(err)
	}

}
