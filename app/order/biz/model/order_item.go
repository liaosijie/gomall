<<<<<<< HEAD
/*
 * @Author: liaosijie
 * @Date: 2025-02-18 16:47:44
 * @Last Modified by: liaosijie
 * @Last Modified time: 2025-02-18 22:14:53
 */

package model

// import (
// 	"gorm.io/gorm"
// 	"time"
// )

// type SnowflakeBase struct {
//     ID        int64 `gorm:"primarykey;autoIncrement:false"`
//     CreatedAt time.Time
//     UpdatedAt time.Time
//     DeletedAt gorm.DeletedAt `gorm:"index"`
// }

type OrderItem struct {
    SnowflakeBase
    ProductID    uint32 `gorm:"type:int(11)"`
    OrderIdRefer string  `gorm:"type:bigint;index"`
    Quantity     uint32 `gorm:"type:int(11)"`
    Cost         float32 `gorm:"type:decimal(10,2)"`
}
=======
/*
 * @Author: liaosijie
 * @Date: 2025-02-18 16:47:44
 * @Last Modified by: liaosijie
 * @Last Modified time: 2025-02-18 22:14:53
 */

package model

type OrderItem struct {
	Base
	ProductID    int64
	OrderIdRefer int64 `gorm:"index"`
	Quantity     int32
	Cost         float32
}
>>>>>>> d44a6b4cc7a74fbb7186470f8fe343cd7e93b530
