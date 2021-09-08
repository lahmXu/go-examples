package example

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/afex/hystrix-go/hystrix"
)

func TestHystrix(t *testing.T) {
	_ = hystrix.Do("wuqq", func() error {
		// talk to other services
		_, err := http.Get("https://www.baidu12.com/")
		if err != nil {
			fmt.Sprintf("get error:%v", err)
			return err
		}
		return nil
	}, func(err error) error {
		fmt.Printf("handle  error:%v\n", err)
		return nil
	})
}
