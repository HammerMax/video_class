package model

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// mongo做底层存储
type kvMongo struct {
	col	*mongo.Collection
}

func (mon *kvMongo) Find(fields map[string]interface{}) (map[string]interface{}, error) {
	cursor, err :=  mon.col.Find(nil, fields)
	if err != nil {
		return nil, err
	}

	result := make(map[string]interface{})
	for cursor.Next(nil) {
		cursor.Decode(&result)
	}

	return result, nil
}

func (mon *kvMongo) Create(fields map[string]interface{}) error {
	_, err := mon.col.InsertOne(nil, fields)
	return err
}

func (mon *kvMongo) Update(query map[string]interface{}, fields map[string]interface{}) error {
	_, err := mon.col.UpdateOne(nil, query, bson.M{"$set": fields})
	return err
}

func (mon *kvMongo) Batch(query []Condition, page, pageSize int) (result []map[string]interface{}, total int, err error) {
	cursor, err := mon.col.Find(nil, parseConditionToFilter(query))
	if err != nil {
		return nil, 0, err
	}

	for cursor.Next(nil) {
	 	single := make(map[string]interface{})
		err := cursor.Decode(&single)
		if err != nil {
			return nil, 0, err
		}

		result = append(result, single)
	}
	return result, 0, nil
}

func parseConditionToFilter(conditions []Condition) bson.D {
	d := bson.D{}
	for _, condition := range conditions {
		switch condition.Mark {
		case EQ:
			d = append(d, bson.E{Key: condition.Key, Value: condition.Value})
		case NEQ:
			d = append(d, bson.E{Key: condition.Key, Value: bson.D{{Key: "$ne", Value: condition.Value}}})
		}
	}
	return d
}

func newKVMongo(addr, db, collection string) (*kvMongo, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(addr))
	if err != nil {
		return nil, err
	}

	err = client.Connect(nil)
	if err != nil {
		return nil, err
	}

	col := client.Database(db).Collection(collection)

	return &kvMongo{col: col}, nil
}