/**
 * @Author: wuchunle<wuchunle@gsaxns.com>
 * @Version: 1.0.0
 * @Description:
 * @File:  orm.go
 * @Time: 2020/9/17 9:43 上午
 */

package orm

import (
	"github.com/987763485/gin-framework/application/conf"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
	"time"
)

func MasterDB() *gorm.DB {
	var db *gorm.DB
	var err error
	db, err = gorm.Open(conf.MDB.Type, conf.MDB.Username+":"+conf.MDB.Username+"@tcp("+conf.MDB.Host+")/"+conf.MDB.DBName+"?charset="+conf.MDB.Charset+"&parseTime=True&loc=Local")
	if err != nil {
		log.Println(err.Error())
		os.Exit(-1) //终止系统
	}
	db.DB().SetMaxIdleConns(10)  //设置闲置的连接数
	db.DB().SetMaxOpenConns(100) //最大打开的连接数，默认值为0表示不限制
	db.DB().SetConnMaxLifetime(time.Second * 3)
	db.SingularTable(true) // 全局禁用表名复数
	db.LogMode(conf.DEBUG)
	return db
}

func SlaveDB() *gorm.DB {
	var db *gorm.DB
	var err error
	db, err = gorm.Open(conf.SDB.Type, conf.SDB.Username+":"+conf.SDB.Password+"@tcp("+conf.SDB.Host+")/"+conf.SDB.DBName+"?charset="+conf.SDB.Charset+"&parseTime=True&loc=Local")
	if err != nil {
		log.Println(err.Error())
		os.Exit(-1) //终止系统
	}
	db.DB().SetMaxIdleConns(10)  //设置闲置的连接数
	db.DB().SetMaxOpenConns(100) //最大打开的连接数，默认值为0表示不限制
	db.DB().SetConnMaxLifetime(time.Second * 3)
	db.SingularTable(true) // 全局禁用表名复数
	db.LogMode(conf.DEBUG)
	return db
}
