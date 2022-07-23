package dsl

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Service struct {
	Db *mongo.Database
}

// Create 创建文档
func (x *Service) Create(ctx context.Context, model string, doc M) (_ interface{}, err error) {
	return x.Db.Collection(model).InsertOne(ctx, doc)
}

// BulkCreate 批量创建文档
func (x *Service) BulkCreate(ctx context.Context, model string, docs []interface{}) (_ interface{}, err error) {
	return x.Db.Collection(model).InsertMany(ctx, docs)
}

// Size 获取文档总数
func (x *Service) Size(ctx context.Context, model string, filter M) (_ int64, err error) {
	if len(filter) == 0 {
		return x.Db.Collection(model).EstimatedDocumentCount(ctx)
	}
	return x.Db.Collection(model).CountDocuments(ctx, filter)
}

// Find 获取匹配文档
func (x *Service) Find(ctx context.Context, model string, filter M, option *options.FindOptions) (data []M, err error) {
	var cursor *mongo.Cursor
	if cursor, err = x.Db.Collection(model).Find(ctx, filter, option); err != nil {
		return
	}
	if err = cursor.All(ctx, &data); err != nil {
		return
	}
	return
}

// FindOne 获取单个文档
func (x *Service) FindOne(ctx context.Context, model string, filter M, option *options.FindOneOptions) (data M, err error) {
	if err = x.Db.Collection(model).FindOne(ctx, filter, option).Decode(&data); err != nil {
		return
	}
	return
}

// Update 局部更新匹配文档
func (x *Service) Update(ctx context.Context, model string, filter M, update M) (_ interface{}, err error) {
	return x.Db.Collection(model).UpdateMany(ctx, filter, update)
}

// UpdateById 局部更新指定 Id 的文档
func (x *Service) UpdateById(ctx context.Context, model string, id primitive.ObjectID, update M) (_ interface{}, err error) {
	return x.Db.Collection(model).UpdateOne(ctx, M{"_id": id}, update)
}

// Replace 替换指定 Id 的文档
func (x *Service) Replace(ctx context.Context, model string, id primitive.ObjectID, doc M) (_ interface{}, err error) {
	return x.Db.Collection(model).ReplaceOne(ctx, M{"_id": id}, doc)
}

// Delete 删除指定 Id 的文档
func (x *Service) Delete(ctx context.Context, model string, id primitive.ObjectID) (_ interface{}, err error) {
	return x.Db.Collection(model).DeleteOne(ctx, M{"_id": id})
}

// BulkDelete 批量删除匹配文档
func (x *Service) BulkDelete(ctx context.Context, model string, filter M) (_ interface{}, err error) {
	return x.Db.Collection(model).DeleteMany(ctx, filter)
}

// Sort 通用排序
func (x *Service) Sort(ctx context.Context, model string, oids []primitive.ObjectID) (_ interface{}, err error) {
	var wms []mongo.WriteModel
	for i, oid := range oids {
		update := M{
			"$set": M{
				"sort":        i,
				"update_time": time.Now(),
			},
		}

		wms = append(wms, mongo.NewUpdateOneModel().
			SetFilter(M{"_id": oid}).
			SetUpdate(update),
		)
	}
	return x.Db.Collection(model).BulkWrite(ctx, wms)
}
