package role

import (
	"net/http"

	"ark-zero-admin/app/core/cmd/api/internal/logic/sys/perm/role"
	"ark-zero-admin/app/core/cmd/api/internal/svc"
	"ark-zero-admin/app/core/cmd/api/internal/types"
	"ark-zero-admin/common/errorx"
	"ark-zero-admin/common/response"

	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateSysRolePermMenuHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateSysRolePermMenuReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, errorx.NewHandlerError(errorx.ParamErrorCode, err.Error()))
			return
		}

		if err := validator.New().StructCtx(r.Context(), req); err != nil {
			httpx.Error(w, errorx.NewHandlerError(errorx.ParamErrorCode, err.Error()))
			return
		}

		l := role.NewUpdateSysRolePermMenuLogic(r.Context(), svcCtx)
		err := l.UpdateSysRolePermMenu(&req)
		if err != nil {
			httpx.Error(w, err)
			return
		}
		response.Response(w, nil, err)
	}
}