package test

import (
	"testing"

	app_context "github.com/holocycle/holo-back/pkg/context2"
	"github.com/holocycle/holo-back/pkg/core"
	"github.com/holocycle/holo-back/pkg/model"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

type ServiceTestcase struct {
	Name           string
	UserID         string
	Precondition   []interface{}
	ExpectCreation []interface{}
	ExpectDeletion []interface{}
	IDGenerator    model.IDGenerator
	Req            []interface{}
	Res            interface{}
	Err            error
}

func DoServiceTest(t *testing.T, testcase ServiceTestcase, sut interface{}) {
	t.Run(testcase.Name, func(t *testing.T) {
		ctx, rollback := GetTestHelper().NewContext(testcase.UserID)
		defer rollback()

		model.SetIDGenerator(testcase.IDGenerator)
		defer model.SetIDGenerator(model.DefaultIDGenerator)

		tx := app_context.GetDB(ctx)
		for _, d := range testcase.Precondition {
			err := tx.Create(d).Error
			if err != nil {
				t.Fatalf("failed to insert precondition: %+v\n", err)
			}
		}

		args := make([]interface{}, 0, len(testcase.Req)+1)
		args = append(args, ctx)
		for i := 0; i < len(testcase.Req); i++ {
			args = append(args, testcase.Req[i])
		}

		res, err := core.Call(sut, args...)
		if err != nil {
			t.Fatalf("failed to call sut: %+v", err)
		}

		if testcase.Err != nil {
			assert.Equal(t, testcase.Err, res[1])
		} else {
			assert.Nil(t, res[1])
		}

		if testcase.Res != nil {
			assert.Equal(t, testcase.Res, res[0])
		} else {
			assert.Nil(t, res[0])
		}

		for _, expected := range testcase.ExpectCreation {
			actual := NewModelSameTypeWith(expected)
			err := tx.Where(expected).Find(actual).Error
			if err != nil && !gorm.IsRecordNotFoundError(err) {
				t.Fatalf("failed to access db: %+v", err)
			}
			assert.NoError(t, err)
		}

		for _, expected := range testcase.ExpectDeletion {
			actual := NewModelSameTypeWith(expected)
			err := tx.Where(expected).Find(actual).Error
			if err != nil && !gorm.IsRecordNotFoundError(err) {
				t.Fatalf("failed to access db: %+v", err)
			}
			assert.Error(t, err)
		}
	})
}

func DoServiceTests(t *testing.T, testcases []ServiceTestcase, sut interface{}) {
	for _, tc := range testcases {
		DoServiceTest(t, tc, sut)
	}
}
