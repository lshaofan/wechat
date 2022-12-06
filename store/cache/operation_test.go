package cache

import "testing"

var redisConfig *RedisConfig

func init() {
	redisConfig = &RedisConfig{
		Host:     "127.0.0.1",
		Port:     63791,
		Prefix:   "wechat:",
		Password: "",
		Database: 0,
	}

}

func TestRedisGet(t *testing.T) {
	ret := NewOperation(redisConfig).Get("testSetNoWithExpire").UnWarp()
	t.Log(ret)

	ret = NewOperation(redisConfig).Get("no_name").UnWarpWithDefault("default")
	t.Log(ret)
}

func TestMGet(t *testing.T) {
	iter := NewOperation(redisConfig).MGet("testSetNoWithExpire", "name1", "name2").Iterator()

	for iter.HasNext() {
		t.Log(iter.Next())
		t.Log("index:", iter.Index)
	}
}

func TestSet(t *testing.T) {

	ret := NewOperation(redisConfig).Set("testSetNoWithExpire", "测试设置，不带过期时间")
	if ret.Result != SetSuccess {
		t.Error("不设置过期时间设置失败：", ret.Err)
	}
}
