package uuidUtil

import (
	"fmt"
	uuid2 "github.com/google/uuid"
	uuid "github.com/satori/go.uuid"
	"github.com/sony/sonyflake"
	"math/rand"
	"sync"
	"time"
)


func GenerateUUID() string {
	u1 := uuid.NewV4()
	fmt.Println(u1)
	return u1.String()
}

var sf *sonyflake.Sonyflake

var rg = struct {
	sync.Mutex
	rand *rand.Rand
}{
	rand: rand.New(rand.NewSource(time.Now().UnixNano())),
}
func init() {
	var f sonyflake.Settings
	f.StartTime = time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	sf = sonyflake.NewSonyflake(f)
	if sf == nil {
		fmt.Println("Init sonyflake err")
	}
}

func Int63nRange(min, max int64) int64 {
	rg.Lock()
	defer rg.Unlock()
	return rg.rand.Int63n(max - min) + min
}

func GenerateSnowflake() string {
	ret, err := sf.NextID()
	if err != nil {
		ret = uint64(Int63nRange(1926425572, 1926425572223607))
	}
	id := fmt.Sprintf("%v", sonyflake.Decompose(ret)["id"])
	return id
}


func GenerateUUID2() {
	// 随机生成一个UUID
	uuid1 := uuid2.New()

	// 基于名称和命名空间生成一个UUID
	uu := uuid2.NewSHA1(uuid2.NamespaceDNS, []byte("example.com"))

	fmt.Println(uuid1)
	fmt.Println(uu)
}
