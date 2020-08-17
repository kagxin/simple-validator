package validator

import (
	"fmt"
	"testing"
)

func Test_Struct(t *testing.T) {

	type Persion struct {
		Name string `gege:"min=1,max=3"`
		Age  int    `gege:"gte=0,lte=12"`
	}
	p := &Persion{Name: "aaaa", Age: 10}
	v := New("gege", ",")
	err := v.Struct(p)
	fmt.Println(err)
}
