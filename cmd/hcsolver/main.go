package main

import (
	"encoding/base64"
	"fmt"

	"github.com/Implex-ltd/hcsolver/cmd/hcsolver/config"
	"github.com/Implex-ltd/hcsolver/cmd/hcsolver/database"
	"github.com/Implex-ltd/hcsolver/cmd/hcsolver/router"
	"github.com/Implex-ltd/hcsolver/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/surrealdb/surrealdb.go"
)

type Fingerprint struct {
	ID          string `json:"id,omitempty"`
	Fingerprint string `json:"fp"`
}

func check_fp() {
	req, err := database.FpDB.Select("fp")
	if err != nil {
		panic(err)
	}

	var FingerprintSlice []Fingerprint
	err = surrealdb.Unmarshal(req, &FingerprintSlice)
	if err != nil {
		panic(err)
	}

	valid := 0
	for _, fp := range FingerprintSlice {
		fmt.Println(fp.ID)
		_, err := base64.RawStdEncoding.DecodeString(fp.Fingerprint)
		if err != nil {
			continue
		}
		valid++

	}

	fmt.Println("valid fp:", valid)
	fmt.Println("total fp:", len(FingerprintSlice))
}

func save_fp() {
	req, err := database.FpDB.Select("fp")
	if err != nil {
		panic(err)
	}

	var FingerprintSlice []Fingerprint
	err = surrealdb.Unmarshal(req, &FingerprintSlice)
	if err != nil {
		panic(err)
	}

	valid := 0
	for _, fp := range FingerprintSlice {
		fmt.Println(fp.ID)
		utils.AppendLine(fp.Fingerprint, "dbfp.txt")
		valid++

	}

	fmt.Println("valid fp:", valid)
	fmt.Println("total fp:", len(FingerprintSlice))
}

func main() {
	// Load the configuration settings
	config.LoadSettings()

	// Connect to the database
	database.ConnectDB(
		config.Config.Database.IP,
		config.Config.Database.Username,
		config.Config.Database.Password,
		config.Config.Database.Port,
	)

	// Print API configuration details
	fmt.Println("API Configuration:")
	fmt.Printf(" - Database IP: %s\n", config.Config.Database.IP)
	fmt.Printf(" - Database Username: %s\n", config.Config.Database.Username)
	fmt.Printf(" - Database Port: %d\n", config.Config.Database.Port)
	fmt.Printf(" - API Port: %d\n", config.Config.API.Port)

	// Perform additional setup or initialization if necessary
	save_fp()

	// Initialize Fiber and set up routes
	app := fiber.New()
	router.SetupRoutes(app)

	// Log that the server is ready and listening
	config.Logger.Info("DB Connected and API online")

	// Start the server and handle errors if they occur
	if err := app.Listen(fmt.Sprintf(":%d", config.Config.API.Port)); err != nil {
		panic(err)
	}
}
