package common

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/go-redis/redis/v8"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

//@author by Hchier
//@Date 2023/1/20 23:29

const ErrLogDest = "logs/err.log"
const VideoDataDest = "static\\video\\data\\"
const VideoCoverDest = "static\\video\\cover\\"
const StaticResources = "http://192.168.0.105:8010/"

func GetFile(dest string) *os.File {
	f, err := os.OpenFile(dest, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err.Error())
	}
	return f
}

var Rdb = GetRdb()

func GetRdb() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func GetRandStr() string {
	result := make([]byte, 16/2)
	println(len(result))
	rand.Read(result)
	return hex.EncodeToString(result)
}

func Log(dest string, v ...interface{}) {
	file := GetFile(dest)
	defer file.Close()
	hlog.SetOutput(file)
	hlog.Error(v...)
}

func ErrLog(v ...interface{}) {
	Log(ErrLogDest, v)
}

// 雪花算法
const (
	workerBits  uint8 = 10                      //机器码位数
	numberBits  uint8 = 12                      //序列号位数
	workerMax   int64 = -1 ^ (-1 << workerBits) //机器码最大值（即1023）
	numberMax   int64 = -1 ^ (-1 << numberBits) //序列号最大值（即4095）
	timeShift   uint8 = workerBits + numberBits //时间戳偏移量
	workerShift uint8 = numberBits              //机器码偏移量
	epoch       int64 = 1656856144640           //起始常量时间戳（毫秒）,此处选取的时间是2022-07-03 21:49:04
)

type Worker struct {
	mu        sync.Mutex
	timeStamp int64
	workerId  int64
	number    int64
}

func NewWorker(workerId int64) (*Worker, error) {
	if workerId < 0 || workerId > workerMax {
		return nil, errors.New("WorkerId超过了限制！")
	}
	return &Worker{
		timeStamp: 0,
		workerId:  workerId,
		number:    0,
	}, nil
}

func (w *Worker) NextId() int64 {
	w.mu.Lock()
	defer w.mu.Unlock()
	//当前时间的毫秒时间戳
	now := time.Now().UnixNano() / 1e6
	//如果时间戳与当前时间相同，则增加序列号
	if w.timeStamp == now {
		w.number++
		//如果序列号超过了最大值，则更新时间戳
		if w.number > numberMax {
			for now <= w.timeStamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else { //如果时间戳与当前时间不同，则直接更新时间戳
		w.number = 0
		w.timeStamp = now
	}
	//ID由时间戳、机器编码、序列号组成
	ID := (now-epoch)<<timeShift | (w.workerId << workerShift) | (w.number)
	return ID
}

func GetSnowId() string {
	worker, err := NewWorker(0)
	if err != nil {
		fmt.Println(err)
	}
	id := worker.NextId()
	return strconv.FormatInt(id, 10)
}

// 验证身份是否有效
func IsValidUser(token string, ctx context.Context) (bool, int64) {
	res, _ := Rdb.HGet(ctx, "tokens", token).Result()
	userId, err := strconv.ParseInt(res, 10, 64)
	if err != nil {
		return false, -1
	}
	return true, userId
}
