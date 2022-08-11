// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	sysRoleFieldNames          = builder.RawFieldNames(&SysRole{})
	sysRoleRows                = strings.Join(sysRoleFieldNames, ",")
	sysRoleRowsExpectAutoSet   = strings.Join(stringx.Remove(sysRoleFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	sysRoleRowsWithPlaceHolder = strings.Join(stringx.Remove(sysRoleFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

	cacheArkAdminZeroSysRoleIdPrefix        = "cache:arkAdminZero:sysRole:id:"
	cacheArkAdminZeroSysRoleUniqueKeyPrefix = "cache:arkAdminZero:sysRole:uniqueKey:"
)

type (
	sysRoleModel interface {
		Insert(ctx context.Context, data *SysRole) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SysRole, error)
		FindOneByUniqueKey(ctx context.Context, uniqueKey string) (*SysRole, error)
		Update(ctx context.Context, data *SysRole) error
		Delete(ctx context.Context, id int64) error
	}

	defaultSysRoleModel struct {
		sqlc.CachedConn
		table string
	}

	SysRole struct {
		Id          int64     `db:"id"`            // 编号
		ParentId    int64     `db:"parent_id"`     // 父级id
		Name        string    `db:"name"`          // 名称
		UniqueKey   string    `db:"unique_key"`    // 唯一标识
		Remark      string    `db:"remark"`        // 备注
		PermMenuIds string    `db:"perm_menu_ids"` // 权限集
		Status      int64     `db:"status"`        // 0=禁用 1=开启
		OrderNum    int64     `db:"order_num"`     // 排序值
		CreateTime  time.Time `db:"create_time"`   // 创建时间
		UpdateTime  time.Time `db:"update_time"`   // 更新时间
	}
)

func newSysRoleModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultSysRoleModel {
	return &defaultSysRoleModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`sys_role`",
	}
}

func (m *defaultSysRoleModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	arkAdminZeroSysRoleIdKey := fmt.Sprintf("%s%v", cacheArkAdminZeroSysRoleIdPrefix, id)
	arkAdminZeroSysRoleUniqueKeyKey := fmt.Sprintf("%s%v", cacheArkAdminZeroSysRoleUniqueKeyPrefix, data.UniqueKey)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, arkAdminZeroSysRoleIdKey, arkAdminZeroSysRoleUniqueKeyKey)
	return err
}

func (m *defaultSysRoleModel) FindOne(ctx context.Context, id int64) (*SysRole, error) {
	arkAdminZeroSysRoleIdKey := fmt.Sprintf("%s%v", cacheArkAdminZeroSysRoleIdPrefix, id)
	var resp SysRole
	err := m.QueryRowCtx(ctx, &resp, arkAdminZeroSysRoleIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysRoleRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSysRoleModel) FindOneByUniqueKey(ctx context.Context, uniqueKey string) (*SysRole, error) {
	arkAdminZeroSysRoleUniqueKeyKey := fmt.Sprintf("%s%v", cacheArkAdminZeroSysRoleUniqueKeyPrefix, uniqueKey)
	var resp SysRole
	err := m.QueryRowIndexCtx(ctx, &resp, arkAdminZeroSysRoleUniqueKeyKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `unique_key` = ? limit 1", sysRoleRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, uniqueKey); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSysRoleModel) Insert(ctx context.Context, data *SysRole) (sql.Result, error) {
	arkAdminZeroSysRoleIdKey := fmt.Sprintf("%s%v", cacheArkAdminZeroSysRoleIdPrefix, data.Id)
	arkAdminZeroSysRoleUniqueKeyKey := fmt.Sprintf("%s%v", cacheArkAdminZeroSysRoleUniqueKeyPrefix, data.UniqueKey)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, sysRoleRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.ParentId, data.Name, data.UniqueKey, data.Remark, data.PermMenuIds, data.Status, data.OrderNum)
	}, arkAdminZeroSysRoleIdKey, arkAdminZeroSysRoleUniqueKeyKey)
	return ret, err
}

func (m *defaultSysRoleModel) Update(ctx context.Context, newData *SysRole) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	arkAdminZeroSysRoleIdKey := fmt.Sprintf("%s%v", cacheArkAdminZeroSysRoleIdPrefix, data.Id)
	arkAdminZeroSysRoleUniqueKeyKey := fmt.Sprintf("%s%v", cacheArkAdminZeroSysRoleUniqueKeyPrefix, data.UniqueKey)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysRoleRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.ParentId, newData.Name, newData.UniqueKey, newData.Remark, newData.PermMenuIds, newData.Status, newData.OrderNum, newData.Id)
	}, arkAdminZeroSysRoleIdKey, arkAdminZeroSysRoleUniqueKeyKey)
	return err
}

func (m *defaultSysRoleModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheArkAdminZeroSysRoleIdPrefix, primary)
}

func (m *defaultSysRoleModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysRoleRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultSysRoleModel) tableName() string {
	return m.table
}
