package menu

import (
	"context"
	"encoding/json"

	"ark-zero-admin/app/core/cmd/api/internal/svc"
	"ark-zero-admin/app/core/cmd/api/internal/types"
	"ark-zero-admin/app/core/model"
	"ark-zero-admin/common/errorx"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSysPermMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSysPermMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSysPermMenuLogic {
	return &UpdateSysPermMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSysPermMenuLogic) UpdateSysPermMenu(req *types.UpdateSysPermMenuReq) error {
	var permMenu = new(model.SysPermMenu)
	err := copier.Copy(permMenu, req)
	if err != nil {
		return errorx.NewDefaultError(errorx.ServerErrorCode)
	}
	bytes, err := json.Marshal(req.Perms)
	if err != nil {
		return errorx.NewDefaultError(errorx.ServerErrorCode)
	}
	permMenu.Perms = string(bytes)
	err = l.svcCtx.SysPermMenuModel.Update(l.ctx, permMenu)
	if err != nil {
		return errorx.NewDefaultError(errorx.ServerErrorCode)
	}
	return nil
}
