package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/ShubhamNatekar/citystore/api/models"
)

var users = []models.User{
	models.User{
		Username: "Shubham Natekar",
		Email:    "shubhamnatekar9221@gmail.com",
		Password: "Shubham@15",
	},
	models.User{
		Username: "Shubham Jagdhane",
		Email:    "shubhamnatekar1010@gmail.com",
		Password: "Shubham@1010",
	},
}

var products = []models.Product{
	models.Product{
		Product_name:  "fastrack M121",
		Cost:3500,
		Quantity:17,
	},
	models.Product{
		Product_name:  "Roadstar F700",
		Cost:900,
		Quantity:8,
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Product{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Product{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Product{}).AddForeignKey("owner_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		products[i].OwnerID = users[i].ID

		err = db.Debug().Model(&models.Product{}).Create(&products[i]).Error
		if err != nil {
			log.Fatalf("cannot seed products table: %v", err)
		}
	}
}

