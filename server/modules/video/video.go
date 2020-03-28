package video

import (
	"reflect"
	"server/modules/model"
)

const tag = "video"

// 视频属性
type Attr struct {
	Vid   string `video:"vid"`
	Title string `video:"title"`
	Image string `video:"image"`
	Cat   int  `video:"cat"`
}

func (attr *Attr) unmarshal(data map[string]interface{}) error {
	t := reflect.TypeOf(*attr)
	v := reflect.ValueOf(attr).Elem()

	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		value, ok := data[sf.Tag.Get(tag)]
		if !ok {
			continue
		}

		v.Field(i).Set(reflect.ValueOf(value))
	}
	return nil
}

func (attr Attr) marshal() map[string]interface{}{
	data := make(map[string]interface{})

	t := reflect.TypeOf(attr)
	v := reflect.ValueOf(attr)
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Tag.Get(tag)] = v.Field(i).Interface()
	}
	return data
}

var engine model.Engine

func InitVideo() {
	var err error
	engine, err = model.NewEngine()
	if err != nil {
		panic(err)
	}
}

func Find(vid string) (*Attr, error) {
	result, err := engine.Find(map[string]interface{}{"vid": vid})
	if err != nil {
		return nil, err
	}

	var a Attr
	err = a.unmarshal(result)
	return &a, err
}

func Create(attr *Attr) error {
	return engine.Create(attr.marshal())
}

func Update(vid string, fields map[string]interface{}) error {
	return engine.Update(map[string]interface{}{"vid": vid}, fields)
}

func List() ([]Attr, error) {
	results, _, err := engine.Batch(nil, 0, 0)
	if err != nil {
		return nil, err
	}

	var as []Attr
	for _, result := range results {
		var a Attr
		err = a.unmarshal(result)
		if err != nil {
			return nil, err
		}
		as = append(as, a)
	}
	return as, nil
}
