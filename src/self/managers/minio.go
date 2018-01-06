/**
 * Created by leeezm on 2017/12/13.
 * Email: shiyi@fightcoder.com
 */

package managers

import (
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	g "self/commons/g"

	"github.com/minio/minio-go"
)

type MinioCli struct {
	cli *minio.Client
}

func NewMinioCli() MinioCli {
	cfg := g.Conf()

	minioClient, err := minio.New(cfg.Minio.Endpoint, cfg.Minio.AccessKeyID, cfg.Minio.SecretAccessKey, cfg.Minio.Secure)
	if err != nil {
		panic(err)
	}

	return MinioCli{cli: minioClient}
}

func (this MinioCli) GetCode(name string) string {
	cfg := g.Conf()
	var flag bool
	resp, err := http.Get("http://xupt1.fightcoder.com:9001/" + cfg.Minio.CodeBucket + "/" + name)
	if err != nil {
		flag = true
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		flag = true
	}
	if flag {
		return ""
	}
	return string(body)
}

func (this MinioCli) GetNameByPath(path string) string {
	strs := strings.Split(path, "/")
	return strs[len(strs)-1]
}

func (this MinioCli) GetImgName(userId int64, picType string) string {
	str := strconv.FormatInt(userId, 10)
	return str + "." + picType
}

func (this MinioCli) GetCodeName() string {
	timestamp := time.Now().UnixNano() / 1000000

	str := strconv.FormatInt(timestamp, 10)
	return str + ".txt"
}

func (this MinioCli) SaveImg(reader io.Reader, userId int64, picType string) string {
	cfg := g.Conf()
	str := this.GetImgName(userId, picType)
	_, err := this.cli.PutObject(cfg.Minio.ImgBucket, str, reader, -1, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		panic(err)
	}
	return str
}

func (this MinioCli) SaveCode(code string) string {
	cfg := g.Conf()
	str := this.GetCodeName()
	_, err := this.cli.PutObject(cfg.Minio.CodeBucket, str, strings.NewReader(code), -1, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		panic(err)
	}
	return str
}

func (this MinioCli) RemoveCode(name string) {
	cfg := g.Conf()
	err := this.cli.RemoveObject(cfg.Minio.CodeBucket, name)
	if err != nil {
		panic(err)
	}
}

func (this MinioCli) Download(objectName, filePath string) {
	cfg := g.Conf()
	err := this.cli.FGetObject(cfg.Minio.CodeBucket, objectName, filePath, minio.GetObjectOptions{})
	if err != nil {
		panic(err)
	}
}
