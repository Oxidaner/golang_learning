package Conns

//
//import (
//	"database/sql/driver"
//	"errors"
//)
//
//func main() {
//	for i := 0; i < maxBadConnRetries; i++ {
//		//从连接池获取连接或通过 driver 新建连接
//		dc, err := db.conn(ctx, strategy)
//		//有空闲连接 -> reuse -> max life time
//		//新建连接 -> max open...
//
//		//将连接放回连接池
//		defer dc.db.putConn(dc, err, true)
//		//validateConnection 有无错误
//		//max life time,max idle conns 检查
//
//		//连接实现 friver.Queryer , driver,Execer 等 interface
//		if err == nil {
//			err = dc.ci.Query(sql, args...)
//		}
//		isBadConn = errors.Is(err, driver.ErrBadConn)
//		if !isBadConn {
//			break
//		}
//	}
//}
