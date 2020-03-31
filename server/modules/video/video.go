package video

import (
	"fmt"
	"reflect"
	"server/modules/model"
	"strconv"
	"time"
)

// 视频属性
type Attr struct {
	Vid   string `json:"vid"`
	Title string `json:"title"`
	Image string `json:"image"`
	CreateTime int `json:"create_time"`
	UpdateTime int `json:"update_time"`
}

func (attr Attr) marshal() map[string]interface{}{
	data := make(map[string]interface{})

	t := reflect.TypeOf(attr)
	v := reflect.ValueOf(attr)
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Tag.Get("json")] = v.Field(i).Interface()
	}
	return data
}

var engine model.Engine

func InitVideo() {
	var err error
	engine, err = model.NewEngine("test")
	if err != nil {
		panic(err)
	}
}

func Find(vid string) (*Attr, error) {
	attr := &Attr{}
	err := engine.Find(map[string]interface{}{"vid": vid}, attr)
	return attr, err
}

func Create(attr *Attr) error {
	attr.CreateTime = int(time.Now().Unix())
	attr.UpdateTime = int(time.Now().Unix())
	return engine.Create(attr.marshal())
}

func Update(vid string, fields map[string]string) error {
	t := reflect.TypeOf(Attr{})
	tags := make(map[string]reflect.Kind)
	for i:=0; i<t.NumField(); i++ {
		sf := t.Field(i)
		tags[sf.Tag.Get("json")] = sf.Type.Kind()
	}

	updateFields := make(map[string]interface{})
	// 默认值
	updateFields["update_time"] = int(time.Now().Unix())

	for key, value := range fields {
		kind, ok := tags[key]
		if !ok {
			return fmt.Errorf("unexpected key:%s", key)
		}

		switch kind {
		case reflect.String:
			updateFields[key] = value
		case reflect.Int:
			iValue, err := strconv.Atoi(value)
			if err != nil {
				return fmt.Errorf("key:%s must be able to convert to int. err:%v", key, err)
			}
			updateFields[key] = iValue

		default:
			return fmt.Errorf("unexpected type. key:%s", key)
		}
	}
	return engine.Update(map[string]interface{}{"vid": vid}, updateFields)
}

func List() ([]Attr, error) {
	var attrs []Attr
	 _, err := engine.Batch(nil, &attrs, 0, 0)
	return attrs, err
}
