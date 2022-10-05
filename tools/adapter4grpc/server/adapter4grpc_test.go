package main

import (
	"fmt"
	"fstools/adapter4grpc/server/service"
	"strconv"
	"testing"
)

func TestXxx(t *testing.T) {
	res := new(service.DirNodeListDto)
	res.Total = 2
	res.Datas = make([]*service.TNode, res.Total)
	for i := int64(0); i < res.Total; i++ {
		t:=new(service.TNode)
		t.Id = strconv.Itoa(int(i))
		res.Datas[i] = t
	}
	fmt.Println(res)
}
