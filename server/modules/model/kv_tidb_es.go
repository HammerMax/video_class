package model

type kvTiEs struct {

}

func (te *kvTiEs) Find(fields map[string]interface{}, result interface{}) error {
	return nil
}

func (te *kvTiEs) Create(fields map[string]interface{}) error {
	return nil
}

func (te *kvTiEs) Update(query map[string]interface{}, fields map[string]interface{}) error {
	return nil
}

func (te *kvTiEs) Batch(query []Condition, result interface{}, page, pageSize int) (total int, err error) {
	return 0, nil
}