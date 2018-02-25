/**
 * Created by leeezm on 2017/12/30.
 * Email: shiyi@fightcoder.com
 */

package managers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/bitly/go-nsq"
	"github.com/pkg/errors"
)

type Nsq struct {
}

type SendMess struct {
	SubmitType string `json:"submit_type"`
	SubmitId   int64  `json:"submit_id"`
}

func (this Nsq) send(topic string, sendMess *SendMess) {
	if topic != "realJudge" && topic != "virtualJudge" {
		err := errors.New("topic is false!")
		panic(err.Error())
	}

	adds := [1]string{"10.0.1.212:9002"}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := r.Intn(len(adds))
	msg, err := json.Marshal(sendMess)
	if err != nil {
		fmt.Println(err)
	}
	postNsq(adds[num], topic, msg)
}

func postNsq(address, topic string, msg []byte) {
	config := nsq.NewConfig()
	if w, err := nsq.NewProducer(address, config); err != nil {
		panic(err)
	} else {
		w.Publish(topic, msg)
	}
}
