package queues

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Controller struct {
	QueuesServices *Service
}

type SyncDto struct {
	Id primitive.ObjectID `json:"id" vd:"mongodb"`
}

func (x *Controller) Sync(ctx context.Context, c *app.RequestContext) {
	var dto SyncDto
	if err := c.BindAndValidate(&dto); err != nil {
		c.Error(err)
		return
	}

	if err := x.QueuesServices.Sync(ctx, dto.Id); err != nil {
		c.Error(err)
		return
	}

	c.Status(204)
}

type DestroyDto struct {
	Ids []primitive.ObjectID `json:"ids" vd:"gt=0,dive,mongodb"`
}

func (x *Controller) Destroy(ctx context.Context, c *app.RequestContext) {
	var dto DestroyDto
	if err := c.BindAndValidate(&dto); err != nil {
		c.Error(err)
		return
	}

	if err := x.QueuesServices.Destroy(ctx, dto.Ids); err != nil {
		c.Error(err)
		return
	}

	c.Status(204)
}

type StateDto struct {
	Id string `path:"id" vd:"mongodb"`
}

func (x *Controller) Info(ctx context.Context, c *app.RequestContext) {
	var dto StateDto
	if err := c.BindAndValidate(&dto); err != nil {
		c.Error(err)
		return
	}

	id, _ := primitive.ObjectIDFromHex(dto.Id)
	r, err := x.QueuesServices.Info(ctx, id)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, r)
}

type PublishDto struct {
	Subject string `json:"subject" vd:"required"`
	Payload M      `json:"payload" vd:"gt=0"`
}

func (x *Controller) Publish(ctx context.Context, c *app.RequestContext) {
	var dto PublishDto
	if err := c.BindAndValidate(&dto); err != nil {
		c.Error(err)
		return
	}

	r, err := x.QueuesServices.Publish(ctx, dto)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, r)
}
