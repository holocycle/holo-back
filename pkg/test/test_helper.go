package test

import (
	"context"

	"github.com/holocycle/holo-back/internal/app/config"
	app_context "github.com/holocycle/holo-back/pkg/context2"
	"github.com/holocycle/holo-back/pkg/db"
	"github.com/holocycle/holo-back/pkg/logger"
	"github.com/holocycle/holo-back/pkg/model"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type TestHelper struct {
	Config *config.AppConfig
	Log    *zap.Logger
	DB     *gorm.DB
}

func InitTestHelper() (func(), error) {
	conf, err := config.NewConfig()
	if err != nil {
		return func() {}, errors.WithStack(errors.WithMessage(err, "Failed to load configuration"))
	}

	log, err := logger.NewLogger(&conf.Logger)
	if err != nil {
		return func() {}, errors.WithStack(errors.WithMessage(err, "Failed to create logger"))
	}

	db, err := db.NewDB(&conf.DB)
	if err != nil {
		log.Fatal("Failed to connect database", zap.Error(err))
	}

	free := func() {
		log.Sync()
		db.Close()
	}

	testHelper = &TestHelper{
		Config: conf,
		Log:    log,
		DB:     db,
	}
	return free, nil
}

func (h *TestHelper) NewContext(userID string) (context.Context, func()) {
	ctx := context.Background()

	session := model.NewSession(userID, nil)
	ctx = app_context.SetSession(ctx, session)

	ctx = app_context.SetLog(ctx, h.Log)

	tx := h.DB.Begin()
	ctx = app_context.SetDB(ctx, tx)
	return ctx, func() {
		tx.Rollback()
	}
}

var testHelper *TestHelper

func GetTestHelper() *TestHelper {
	return testHelper
}
