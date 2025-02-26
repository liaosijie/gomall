<<<<<<< HEAD
/*
 * @Author: liaosijie
 * @Date: 2025-02-18 16:47:32
 * @Last Modified by: liaosijie
 * @Last Modified time: 2025-02-18 17:08:47
 */

package mysql

import (
	//"douyin-gomall/gomall/app/order/biz/model"
	//"douyin-gomall/gomall/app/order/conf"
	"fmt"
	"os"

	"github.com/PiaoAdmin/gomall/app/order/biz/model"
	"github.com/PiaoAdmin/gomall/app/order/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"k8s.io/klog/v2"
=======
package mysql

import (
	"fmt"
	"os"

	"github.com/PiaoAdmin/gomall/app/order/biz/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
>>>>>>> b6e73c27fce12b01552c5334097a847176b8f26a
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
<<<<<<< HEAD
<<<<<<< HEAD
	dsn :=fmt.Sprint(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	DB, err = gorm.Open(mysql.Open(dsn),
=======
	DB, err = gorm.Open(mysql.Open(conf.GetConf().MySQL.DSN),
>>>>>>> b6e73c27fce12b01552c5334097a847176b8f26a
=======
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/user?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	// dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))

	DB, err = gorm.Open(mysql.Open(dsn),
>>>>>>> d44a6b4cc7a74fbb7186470f8fe343cd7e93b530
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
<<<<<<< HEAD
<<<<<<< HEAD
	if os.Getenv("GO_ENV") != "online"{
		if err := DB.AutoMigrate(&model.Order{},&model.OrderItem{}); err!= nil {
			klog.Error(err)
		}
	}
=======
>>>>>>> b6e73c27fce12b01552c5334097a847176b8f26a
=======
	if os.Getenv("GO_ENV") != "online" {
		DB.AutoMigrate(
			&model.Order{},
			&model.OrderItem{},
		)
		DB = DB.Debug()
	}
>>>>>>> d44a6b4cc7a74fbb7186470f8fe343cd7e93b530
}
