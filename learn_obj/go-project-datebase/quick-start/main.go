package quick_start

import (
	"database/sql"
)

type User struct {
	ID   int
	Name string
}

func main() {
	//使用driver + DSN 初始化DB连接
	db, err := sql.Open("mysql", "zzb:123456@tcp(127.0.0.1:3306)/hello ")

	//执行一条sql.通过rows取回返回的数据处理完毕,需要释放连接
	rows, err := db.Query("select id ,name from users where id = ?", 1)
	if err != nil {
		//XXX
	}

	//数据错误处理
	defer rows.Close() //可能会返回一个错误,需要进行检查
	defer func() {
		err = rows.Close()
	}()

	var users []User
	for rows.Next() {
		var user User
		rows.Scan(&user.ID, &user.Name)
		if err != nil {

		}

		users = append(users, user)
		//处理错误
		if rows.Err() != nil {
			//....
		}
	}
}
