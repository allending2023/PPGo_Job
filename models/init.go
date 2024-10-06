/*
* @Author: haodaquan
* @Date:   2017-06-20 09:44:44
* @Last Modified by:   Bee
* @Last Modified time: 2019-02-15 22:12
 */

package models

import (
	"fmt"
	"net/url" // 标准库中的fmt包来进行控制台输出

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

var StartTime int64

func Init(startTime int64) {
	StartTime = startTime
	dbType := beego.AppConfig.String("db.type")
	if dbType == "" {
		dbType = "mysql"
	}
	dbhost := beego.AppConfig.String("db.host")
	dbport := beego.AppConfig.String("db.port")
	dbuser := beego.AppConfig.String("db.user")
	dbpassword := beego.AppConfig.String("db.password")
	dbname := beego.AppConfig.String("db.name")
	timezone := beego.AppConfig.String("db.timezone")
	if dbport == "" {
		if dbType == "mysql" {
			dbport = "3306"
		} else if dbType == "postgres" {
			dbport = "5432"
		}
	}

	dsn := ""
	// 注册驱动
	if dbType == "mysql" {
		orm.RegisterDriver("postgres", orm.DRMySQL)
		dsn = dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
		if timezone != "" {
			dsn = dsn + "&loc=" + url.QueryEscape(timezone)
		}
		orm.RegisterDataBase("default", dbType, dsn)
	} else if dbType == "postgres" {
		orm.RegisterDriver("postgres", orm.DRPostgres)
		dsn = "user=" + dbuser + " password=" + dbpassword + " dbname=" + dbname + " host=" + dbhost + " port=" + dbport + " sslmode=disable"
		orm.RegisterDataBase("default", dbType, dsn)
	}

	fmt.Println("init ...dbType=" + dbType + ", dsn=" + dsn)

	orm.RegisterModel(
		new(Admin),
		new(Auth),
		new(Role),
		new(RoleAuth),
		new(ServerGroup),
		new(TaskServer),
		new(Ban),
		new(Group),
		new(Task),
		new(TaskLog),
		new(NotifyTpl),
	)

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
}

func TableName(name string) string {
	return beego.AppConfig.String("db.prefix") + name
}
