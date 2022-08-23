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
	sysLogFieldNames          = builder.RawFieldNames(&SysLog{})
	sysLogRows                = strings.Join(sysLogFieldNames, ",")
	sysLogRowsExpectAutoSet   = strings.Join(stringx.Remove(sysLogFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	sysLogRowsWithPlaceHolder = strings.Join(stringx.Remove(sysLogFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

	cacheArkAdminSysLogIdPrefix = "cache:arkAdmin:sysLog:id:"
)

type (
	sysLogModel interface {
		Insert(ctx context.Context, data *SysLog) (sql.Result, error)
		FindOne(ctx context.Context, id uint64) (*SysLog, error)
		Update(ctx context.Context, data *SysLog) error
		Delete(ctx context.Context, id uint64) error
	}

	defaultSysLogModel struct {
		sqlc.CachedConn
		table string
	}

	SysLog struct {
		Id         uint64    `db:"id"`          // 编号
		UserId     int64     `db:"user_id"`     // 操作账号
		Ip         string    `db:"ip"`          // ip
		Uri        string    `db:"uri"`         // 请求路径
		Type       int64     `db:"type"`        // 1=登录日志 2=操作日志
		Request    string    `db:"request"`     // 请求数据
		Status     int64     `db:"status"`      // 0=失败 1=成功
		CreateTime time.Time `db:"create_time"` // 创建时间
		UpdateTime time.Time `db:"update_time"` // 更新时间
	}
)

func newSysLogModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultSysLogModel {
	return &defaultSysLogModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`sys_log`",
	}
}

func (m *defaultSysLogModel) Delete(ctx context.Context, id uint64) error {
	arkAdminSysLogIdKey := fmt.Sprintf("%s%v", cacheArkAdminSysLogIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, arkAdminSysLogIdKey)
	return err
}

func (m *defaultSysLogModel) FindOne(ctx context.Context, id uint64) (*SysLog, error) {
	arkAdminSysLogIdKey := fmt.Sprintf("%s%v", cacheArkAdminSysLogIdPrefix, id)
	var resp SysLog
	err := m.QueryRowCtx(ctx, &resp, arkAdminSysLogIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysLogRows, m.table)
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

func (m *defaultSysLogModel) Insert(ctx context.Context, data *SysLog) (sql.Result, error) {
	arkAdminSysLogIdKey := fmt.Sprintf("%s%v", cacheArkAdminSysLogIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, sysLogRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.Ip, data.Uri, data.Type, data.Request, data.Status)
	}, arkAdminSysLogIdKey)
	return ret, err
}

func (m *defaultSysLogModel) Update(ctx context.Context, data *SysLog) error {
	arkAdminSysLogIdKey := fmt.Sprintf("%s%v", cacheArkAdminSysLogIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysLogRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserId, data.Ip, data.Uri, data.Type, data.Request, data.Status, data.Id)
	}, arkAdminSysLogIdKey)
	return err
}

func (m *defaultSysLogModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheArkAdminSysLogIdPrefix, primary)
}

func (m *defaultSysLogModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysLogRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultSysLogModel) tableName() string {
	return m.table
}
