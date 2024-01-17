package injector

import (
	"github.com/ryo034/react-go-template/apps/system/api/interface/gateway/me"
	"github.com/ryo034/react-go-template/apps/system/api/interface/gateway/member"
	"github.com/ryo034/react-go-template/apps/system/api/interface/gateway/user"
)

type GatewayAdapter struct {
	User   user.Adapter
	Member member.Adapter
	Me     me.Adapter
}

func newGatewayAdapterInjector() GatewayAdapter {
	userGw := user.NewAdapter()
	memberGw := member.NewAdapter(userGw)
	meGw := me.NewAdapter(userGw, memberGw)
	return GatewayAdapter{
		userGw, memberGw, meGw,
	}
}
