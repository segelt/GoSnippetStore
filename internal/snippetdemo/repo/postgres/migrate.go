package snippetrepo

import (
	"fmt"
	client "snippetdemo/internal/database/postgres"
	"snippetdemo/pkg/models"
)

func MigrateModels() {
	err := client.DbClient.AutoMigrate(&models.Snippet{})

	if err != nil {
		panic(err)
	}
	err = client.DbClient.AutoMigrate(&models.User{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Migrated models")
}
