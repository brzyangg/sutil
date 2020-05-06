package statlog

import (
	"context"
	"github.com/shawnfeng/sutil/scontext"
	"testing"
)

type sa struct {
	s string
	i int
}

type testHead struct {
	uid     int64
	source  int32
	ip      string
	region  string
	dt      int32
	unionid string
}

func (th *testHead) ToKV() map[string]interface{} {
	return map[string]interface{}{
		"uid":     th.uid,
		"source":  th.source,
		"ip":      th.ip,
		"region":  th.region,
		"dt":      th.dt,
		"unionid": th.unionid,
	}
}

func TestLog(t *testing.T) {
	Init("", "", "testservice")
	t.Run("context without head", func(t *testing.T) {
		LogKV(context.TODO(), "method1",
			"k1", 0,
			"k2", "hello",
			"k3", true,
			"k4", []int{1, 2, 3},
			"k5", &sa{"world", 0})
	})

	t.Run("context with head", func(t *testing.T) {
		ctx := context.TODO()
		ctx = context.WithValue(ctx, scontext.ContextKeyHead, &testHead{
			uid:     1234,
			source:  0,
			ip:      "192.168.0.1",
			region:  "asia",
			dt:      0,
			unionid: "unionid",
		})
		LogKV(ctx, "method2", "k1", 0,
			"k2", "hello",
			"k3", true,
			"k4", []int{1, 2, 3},
			"k5", &sa{"world", 0})
	})

	t.Run("ignore non-equal keys and values", func(t *testing.T) {
		LogKV(context.TODO(), "method3", "k1", 0,
			"k2", "hello",
			"k3", true,
			"k4", []int{1, 2, 3},
			"k5")
	})

	t.Run("fail-fast if key is not of string type", func(t *testing.T) {
		LogKV(context.TODO(), "method4",
			0, "k1",
			"k2", 1)
	})
}

func TestLogKVWithMap(t *testing.T) {
	Init("", "", "testservice")
	t.Run("context without head", func(t *testing.T) {
		LogKVWithMap(context.TODO(), "method1", map[string]interface{}{
			"k1": 0,
			"k2": "hello",
			"k3": true,
			"k4": []int{1, 2, 3},
			"k5": &sa{"world", 0},
		})
	})

	t.Run("context without head LogKV", func(t *testing.T) {
		LogKV(context.TODO(), "method1",
			"k1", 0,
			"k2", "hello",
			"k3", true,
			"k4", []int{1, 2, 3},
			"k5", &sa{"world", 0})
	})

	t.Run("context with head", func(t *testing.T) {
		ctx := context.TODO()
		ctx = context.WithValue(ctx, scontext.ContextKeyHead, &testHead{
			uid:     1234,
			source:  0,
			ip:      "192.168.0.1",
			region:  "asia",
			dt:      0,
			unionid: "unionid",
		})
		LogKVWithMap(ctx, "method2", map[string]interface{}{
			"k1": 0,
			"k2": "hello",
			"k3": true,
			"k4": []int{1, 2, 3},
			"k5": &sa{"world", 0},
		})
	})
}
