package test

import (
	"context"
	"testing"

	app_context "github.com/holocycle/holo-back/pkg/context2"
	"github.com/holocycle/holo-back/pkg/model"
	"github.com/stretchr/testify/assert"
)

type Service func(ctx context.Context, req interface{}) (res interface{}, err error)

type ServiceTestcase struct {
	Name          string
	UserID        string
	Precondition  []interface{}
	Postcondition []interface{}
	IDGenerator   model.IDGenerator
	Req           interface{}
	Res           interface{}
	Err           error
}

func DoServiceTest(t *testing.T, testcase ServiceTestcase, sut Service) {
	t.Run(testcase.Name, func(t *testing.T) {
		ctx, rollback := GetTestHelper().NewContext(testcase.UserID)
		defer rollback()

		model.SetIDGenerator(testcase.IDGenerator)
		defer model.SetIDGenerator(model.DefaultIDGenerator)

		tx := app_context.GetDB(ctx)
		for _, d := range testcase.Precondition {
			err := tx.Create(d).Error
			if err != nil {
				t.Fatalf("%+v\n", err)
			}
		}

		res, err := sut(ctx, testcase.Req)
		if testcase.Err != nil {
			assert.Equal(t, testcase.Err, err)
		} else {
			assert.Nil(t, err)
		}

		if testcase.Res != nil {
			assert.Equal(t, testcase.Res, res)
		} else {
			assert.Nil(t, res)
		}

		for _, expected := range testcase.Postcondition {
			actual := NewModelSameTypeWith(expected)
			err := tx.Where(expected).Find(actual).Error
			assert.NoError(t, err)
		}
	})
}

func DoServiceTests(t *testing.T, testcases []ServiceTestcase, sut Service) {
	for _, tc := range testcases {
		DoServiceTest(t, tc, sut)
	}
}
