package test

import (
	"bytes"
	"cloud-disk/model"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"xorm.io/xorm"
)

func TestXormTest(t *testing.T) {
	engine, err := xorm.NewEngine("mysql", "root:root@tcp(192.168.179.142:3306)/cloud-disk?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		t.Fatal(err)
	}
	data := make([]model.UserBasic, 0)
	err = engine.Find(&data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("this is data", data)
	b, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("this is b ", string(b))
	dst := new(bytes.Buffer)
	err = json.Indent(dst, b, "", "")
	if err != nil {
		t.Fatal(err)
	}
	var res []model.UserBasic
	err = json.Unmarshal(b, &res)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v", res)
	//fmt.Println(dst)
}
