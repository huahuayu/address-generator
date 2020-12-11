package session

import (
	"encoding/json"
	"fmt"
	"github.com/huahuayu/address-generator/common/redis"
	"github.com/huahuayu/address-generator/global"
	"github.com/huahuayu/address-generator/model"
)

func Set(sid string, user *model.TUser) {
	key := fmt.Sprintf(redis.KEY_USER_SESSION, sid)
	data, _ := json.Marshal(user)
	redis.Client.Set(key, data, global.SessionExpiredTime)
}

func Get(sid string) *model.TUser {
	key := fmt.Sprintf(redis.KEY_USER_SESSION, sid)
	res := redis.Client.Get(key)
	user := &model.TUser{}
	bytes, _ := res.Bytes()
	err := json.Unmarshal(bytes, user)
	if err != nil {
		return nil
	}
	return user
}

func Del(sid string) {
	key := fmt.Sprintf(redis.KEY_USER_SESSION, sid)
	redis.Client.Del(key)
}
