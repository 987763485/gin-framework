/**
 * @Author: wuchunle<wuchunle@gsaxns.com>
 * @Version: 1.0.0
 * @Description:
 * @File:  conf
 * @Time: 2020/9/17 9:56 上午
 */

package conf

type dbConf struct {
	Type     string
	Host     string
	Username string
	Password string
	DBName   string
	Charset  string
}

const DEBUG = true

var MDB = dbConf{
	Type:     "mysql",
	Host:     "127.0.0.1",
	Username: "root",
	Password: "root",
	DBName:   "db_name",
	Charset:  "utf8mb4",
}

var SDB = dbConf{
	Type:     "mysql",
	Host:     "127.0.0.1",
	Username: "root",
	Password: "root",
	DBName:   "db_name",
	Charset:  "utf8mb4",
}
