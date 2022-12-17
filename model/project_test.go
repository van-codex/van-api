package model_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/weplanx/server/model"
	"testing"
)

func TestCreateProject(t *testing.T) {
	var err error
	_, err = db.Collection("projects").InsertOne(
		context.TODO(),
		model.NewProject("DEFAULT", "default"),
	)
	assert.NoError(t, err)
}