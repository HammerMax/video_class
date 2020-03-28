package model

type Condition struct {
	Mark int
	Key string
	Value interface{}
}

const (
	EQ = iota // 等于
	NEQ       // 不等于
	GT        // 大于
	LT 		  // 小于
	LIKE	  // 匹配
)

// 视频属性
type Attr struct {
	Vid   string
	Title string
	Image string
	Cat   int
}

// 定义视频数据类型和操作方法
type Engine interface {
	Find(fields map[string]interface{}) (map[string]interface{}, error)
	Create(fields map[string]interface{}) error
	Update(query map[string]interface{}, fields map[string]interface{}) error

	Batch(query []Condition, page, pageSize int) (result []map[string]interface{}, total int, err error)
}

func NewEngine() (Engine, error) {
	return newKVMongo("mongodb://9.134.33.32:27017", "test", "test")
}
