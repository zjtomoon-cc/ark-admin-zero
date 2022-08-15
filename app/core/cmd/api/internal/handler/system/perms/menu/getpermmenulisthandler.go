package menu

import (
	"net/http"

	"ark-zero-admin/app/core/cmd/api/internal/logic/system/perms/menu"
	"ark-zero-admin/app/core/cmd/api/internal/svc"
	"ark-zero-admin/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetPermMenuListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := menu.NewGetPermMenuListLogic(r.Context(), svcCtx)
		resp, err := l.GetPermMenuList()
		if err != nil {
			httpx.Error(w, err)
			return
		}
		response.Response(w, resp, err)
	}
}