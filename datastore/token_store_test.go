package datastore

import (
	"testing"

	"github.com/garyburd/redigo/redis"
	"github.com/patelndipen/AP1/settings"
)

var GlobalTokenStore *JWTStore

func init() {
	settings.SetPreproductionEnv()
	GlobalTokenStore = &JWTStore{ConnectToRedis()}
}

func TestStoreToken(t *testing.T) {

	key := "key"
	val := "val"

	GlobalTokenStore.StoreToken(key, val, 100)

	retrievedVal, err := redis.String(GlobalTokenStore.Conn.Do("GET", key))
	if err != nil {
		t.Error(err)
	}

	if retrievedVal != val {
		t.Errorf("The retrieved value: %s, is not the same as the value stored in the redis: %s", retrievedVal, val)
	}

}

func TestIsTokenStored(t *testing.T) {

	if !GlobalTokenStore.IsTokenStored("key") {
		t.Errorf("IsTokenStore did not recognize that \"key\" is stored in redis")
	}
}
