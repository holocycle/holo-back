package controller

import (
	"net/http"

	"github.com/holocycle/holo-back/internal/app/config"
	app_context "github.com/holocycle/holo-back/internal/app/context"
	"github.com/holocycle/holo-back/pkg/api"
	"github.com/holocycle/holo-back/pkg/context"
	app_context2 "github.com/holocycle/holo-back/pkg/context2"
	"github.com/holocycle/holo-back/pkg/converter"
	"github.com/holocycle/holo-back/pkg/model"
	"github.com/holocycle/holo-back/pkg/repository"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type CommentController struct {
	Config            *config.AppConfig
	CommentRepository repository.CommentRepository
}

func NewCommentController(config *config.AppConfig) *CommentController {
	return &CommentController{
		Config:            config,
		CommentRepository: repository.NewCommentRepository(),
	}
}

func (c *CommentController) Register(e *echo.Echo) {
	get(e, "/clips/:clip_id/comments", c.ListComments)
	get(e, "/clips/:clip_id/comments/:comment_id", c.GetComment)
	post(e, "/clips/:clip_id/comments", c.PostComment)
	delete(e, "/clips/:clip_id/comments/:comment_id", c.DeleteComment)
}

func (c *CommentController) ListComments(ctx context.Context) error {
	log := ctx.GetLog()
	goCtx := app_context2.FromEchoContext(ctx)

	clipID := ctx.Param("clip_id")
	if clipID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "please specify clip_id")
	}
	log.Info("success to retrieve path parameter", zap.String("clipId", clipID))

	req := &api.ListCommentsRequest{}
	if err := inject(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	log.Info("success to validate", zap.Any("req", req))

	// TODO: clipIDが実在することのバリデーション処理

	query := c.CommentRepository.NewQuery(goCtx)
	if req.Limit > 0 {
		query = query.Limit(req.Limit)
	}
	if req.OrderBy == "latest" {
		query = query.Latest()
	}
	comments, err := query.
		Where(&model.Comment{ClipID: clipID}).
		JoinUser().
		FindAll()
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, &api.ListCommentsResponse{
		Comments: converter.ConvertToComments(comments),
	})
}

func (c *CommentController) GetComment(ctx context.Context) error {
	log := ctx.GetLog()
	goCtx := app_context2.FromEchoContext(ctx)

	clipID := ctx.Param("clip_id")
	if clipID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "please specify clip_id")
	}
	commentID := ctx.Param("comment_id")
	if commentID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "please specify comment_id")
	}
	log.Info("success to retrieve path parameter", zap.String("clipId", clipID), zap.String("commentId", commentID))

	req := &api.GetCommentRequest{}
	if err := inject(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	log.Info("success to validate", zap.Any("req", req))

	// TODO: clipIDが実在することのバリデーション処理

	comment, err := c.CommentRepository.NewQuery(goCtx).
		Where(
			&model.Comment{
				ID:     commentID,
				ClipID: clipID,
			}).
		JoinUser().
		Find()
	if err != nil {
		if repository.NotFoundError(err) {
			return echo.NewHTTPError(http.StatusNotFound, "comment was not found")
		}
		return err
	}

	return ctx.JSON(http.StatusOK, &api.GetCommentResponse{
		Comment: converter.ConvertToComment(comment),
	})
}

func (c *CommentController) PostComment(ctx context.Context) error {
	log := ctx.GetLog()
	goCtx := app_context2.FromEchoContext(ctx)

	req := &api.PostCommentRequest{}
	if err := inject(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	clipID := ctx.Param("clip_id")
	if clipID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "please specify clip_id")
	}
	log.Debug("success to validate request", zap.String("clipID", clipID))

	// TODO: clipIDが実在することのバリデーション処理

	comment := model.NewComment(
		app_context.GetSession(ctx).UserID,
		clipID,
		req.Content,
	)
	if err := c.CommentRepository.NewQuery(goCtx).Create(comment); err != nil {
		log.Error("failed to create comment", zap.Any("comment", comment))
		return err
	}
	log.Info("success to create comment", zap.Any("comment", comment))

	return ctx.JSON(http.StatusCreated, &api.PostCommentResponse{
		CommentID: comment.ID,
	})
}

func (c *CommentController) DeleteComment(ctx context.Context) error {
	log := ctx.GetLog()
	goCtx := app_context2.FromEchoContext(ctx)

	req := &api.DeleteCommentRequest{}
	if err := inject(ctx, req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	clipID := ctx.Param("clip_id")
	if clipID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "please specify clip_id")
	}
	commentID := ctx.Param("comment_id")
	if commentID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "please specify comment_id")
	}
	log.Debug("success to validate request", zap.String("clipID", clipID), zap.String("commentID", commentID))

	// TODO: clipIDが実在することのバリデーション処理

	cond := &model.Comment{
		ID:     commentID,
		UserID: app_context.GetSession(ctx).UserID,
		ClipID: clipID,
	}
	rows, err := c.CommentRepository.NewQuery(goCtx).Where(cond).Delete()
	if err != nil {
		return err
	}
	if rows == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "comment couldn't found")
	}

	return ctx.JSON(http.StatusOK, &api.DeleteCommentResponse{})
}
