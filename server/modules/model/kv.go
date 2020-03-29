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

// 定义视频数据类型和操作方法
type Engine interface {
	Find(fields map[string]interface{}, result interface{}) error
	Create(fields map[string]interface{}) error
	Update(query map[string]interface{}, fields map[string]interface{}) error

	Batch(query []Condition, result interface{}, page, pageSize int) (total int, err error)
}

func NewEngine() (Engine, error) {
	return newKVMongo("mongodb://9.134.33.32:27017", "test", "test")
}
