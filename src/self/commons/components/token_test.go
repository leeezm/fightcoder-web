/**
 * Created by leeezm on 2017/12/27.
 * Email: shiyi@fightcoder.com
 */

package components

import (
	"fmt"
	"testing"
	"time"
)

func Test(t *testing.T) {
	head := &HeaderData{"HS256", "JWT"}
	pay := &PayLoadData{"1514618995229", "John Doe", "12"}
	token := GetToken(head, pay)
	fmt.Println("token-> ", token)

	fmt.Println("res-> ", CheckToken(token))
	fmt.Println("PayLoadData:", GetPayLoad(token))
}

func TestOne(t *testing.T) {
	fmt.Println(time.Now().UnixNano() / 1000000)
}
