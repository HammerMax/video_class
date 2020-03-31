package file_metadata

import (
	"reflect"
	"server/modules/model"
	"time"
)

const DataDir = "./tmp"

// 文件元数据
type Metadata struct {
	Key string `json:"key"`
	Name string `json:"name"`
	Size int64    `json:"size"`
	CreateTime int `json:"create_time"`
	UpdateTime int `json:"update_time"`
}

func (md Metadata) marshal() map[string]interface{}{
	data := make(map[string]interface{})

	t := reflect.TypeOf(md)
	v := reflect.ValueOf(md)
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Tag.Get("json")] = v.Field(i).Interface()
	}
	return data
}

var engine model.Engine

func InitEngine() {
	var err error
	engine, err = model.NewEngine("storage")
	if err != nil {
		panic(err)
	}
}

func Create(m *Metadata) error {
	m.CreateTime = int(time.Now().Unix())
	m.UpdateTime = int(time.Now().Unix())
	return engine.Create(m.marshal())
}
