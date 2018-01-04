/**
 * Created by shiyi on 2017/10/1.
 * Email: shiyi@fightcoder.com
 */

package store

func InitStore() {
	InitMysql()
}

func CloseStore() {
	closeMysql()
}
