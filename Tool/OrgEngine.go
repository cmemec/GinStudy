package tool

import (
	model "GinStudy/Model"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var XormEngine *xorm.Engine

func OrmEngine(cfg *Config) error {
	log.Println("start:connect mysql")
	database := cfg.Database
	conn := database.User + ":" + database.Password + "@tcp(" + database.Host + ":" + database.Port + ")/" + database.DbName + "?charset=" + database.Charset
	log.Println(database)
	engine, err := xorm.NewEngine(database.Driver, conn)
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Println("start:connect success")
	engine.ShowSQL(database.ShowSql)
	err = engine.Sync2(new(model.SmsCode), new(model.User))
	if err != nil {
		return err
	}
	XormEngine = engine
	return nil
}
