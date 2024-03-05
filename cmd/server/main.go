package main

import (
	router "fiber-tutorial/api/rest/routes"
	"fiber-tutorial/internal/config"
	"fiber-tutorial/internal/datastore"
	"fiber-tutorial/pkg/i18n"
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pressly/goose"
	"golang.org/x/text/language"
)

func main() {

	cfgPath := "internal/config/config_dev.yaml"
	cfg, err := config.NewConfig(cfgPath)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	InitTranslator()
	mysqlDatabaseMigrate(cfg)

	db, err := datastore.NewMysqldb()
	if err != nil {
		fmt.Println("error wihle connect database")
	}

	db.InitTables()

	app := fiber.New()
	app.Use(cors.New(
		cors.Config{
			AllowOrigins:     strings.Join(cfg.Http.AllowedOrigins, ","),
			AllowMethods:     strings.Join(cfg.Http.AllowedMethods, ","),
			AllowHeaders:     strings.Join(cfg.Http.AllowedHeaders, ","),
			AllowCredentials: cfg.Http.AllowCredentials,
			ExposeHeaders:    strings.Join(cfg.Http.ExposedHeaders, ","),
			MaxAge:           cfg.Http.MaxAge,
		},
	))

	router.InitRoutes(app, db)

	app.Listen(cfg.Http.Port)

}

func InitTranslator() {
	translator := i18n.NewTranslator(language.English)
	err := translator.LoadLanguageFile("internal/model/translate/en.json")
	if err != nil {
		fmt.Printf("err: %s", err)
	}
	err = translator.LoadLanguageFile("internal/model/translate/tr.json")
	if err != nil {
		fmt.Printf("err: %s", err)
	}
	i18n.Translator = translator
}

// bakılacak
func mysqlDatabaseMigrate(cfg *config.Config) error {
	db, err := goose.OpenDBWithDriver("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&multiStatements=true", cfg.Mysql_Db.User, cfg.Mysql_Db.Password, cfg.Mysql_Db.Host, cfg.Mysql_Db.Port, cfg.Mysql_Db.Dbname))
	if err != nil {
		fmt.Printf("err: %s", err)
		return err
	}
	if err := goose.Up(db, cfg.Mysql_Db.MigrationPath); err != nil {
		fmt.Printf("err: %s", err)
		return err
	}
	if err := db.Close(); err != nil {
		fmt.Printf("err: %s", err)
		return err
	}
	return nil
}
