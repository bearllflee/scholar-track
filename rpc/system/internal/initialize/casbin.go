package initialize

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
	"sync"
)

var (
	syncedCachedEnforcer *casbin.SyncedCachedEnforcer
	once                 sync.Once
)

func Casbin(db *gorm.DB) *casbin.SyncedCachedEnforcer {
	once.Do(func() {
		adapter, err := gormadapter.NewAdapterByDB(db)
		if err != nil {
			return
		}
		text := `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
		`
		m, err := model.NewModelFromString(text)
		if err != nil {
			return
		}
		syncedCachedEnforcer, _ = casbin.NewSyncedCachedEnforcer(m, adapter)
		syncedCachedEnforcer.SetExpireTime(60 * 60)
		err = syncedCachedEnforcer.LoadPolicy()
		if err != nil {
			return
		}
	})
	return syncedCachedEnforcer
}
