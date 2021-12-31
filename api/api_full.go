package api

import "github.com/filecoin-project/venus-wallet/api/permission"

var _ IFullAPI = &FullAPIAdapter{}

// full service API permissions constraints
type FullAPIAdapter struct {
	CommonAPIAdapter
	WalletAPIAdapter
	WalletLockAPIAdapter
	StrategyAPIAdapter
	WalletEventAPIAdapter
}

func PermissionedFullAPI(a IFullAPI) IFullAPI {
	var out FullAPIAdapter
	permission.PermissionedAny(a, &out.WalletAPIAdapter.Internal)
	permission.PermissionedAny(a, &out.CommonAPIAdapter.Internal)
	permission.PermissionedAny(a, &out.WalletLockAPIAdapter.Internal)
	permission.PermissionedAny(a, &out.StrategyAPIAdapter.Internal)
	permission.PermissionedAny(a, &out.WalletEventAPIAdapter.Internal)
	return &out
}
