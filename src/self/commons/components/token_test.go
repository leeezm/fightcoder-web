/**
 * Created by leeezm on 2017/12/27.
 * Email: shiyi@fightcoder.com
 */

package components

import (
	"encoding/base64"
	"fmt"
	"testing"
	"time"
)

func Test(t *testing.T) {
	head := &HeaderData{"HS256", "JWT"}
	pay := &PayLoadData{"1514618995229", "John Doe", "12"}
	token := GetToken(head, pay)
	fmt.Println("token-> ", token)

	fmt.Println("res-> ", CheckToken("SFMyNTYuSldU.MTUxNTM0MjExNzQxOS7mtYvor5UuNQ==.0xWosreVMI+/C1Vz1DCkLhzuicVbQob3dEGDSkzdp6E="))
	fmt.Println("PayLoadData:", GetPayLoad(token))
}

func TestOne(t *testing.T) {
	urltest := "SFMyNTYuSldU.MTUxNTM0MjExNzQxOS7mtYvor5UuNQ==.0xWosreVMI+/C1Vz1DCkLhzuicVbQob3dEGDSkzdp6E="

	fmt.Println(urltest)

	encodeurl := base64.StdEncoding.EncodeToString([]byte(urltest))

	fmt.Println(encodeurl)

	//a, _ := url.QueryUnescape(encodeurl)
	//fmt.Println(a)

	fmt.Println(time.Now().UnixNano() / 1000000)
}
