package app

import (
	sdk "github.com/Null-nil/cosmos-sdk/types"
	"github.com/Null-nil/cosmos-sdk/types/module"
	upgradetypes "github.com/Null-nil/cosmos-sdk/x/upgrade/types"
)

func (app *EthermintApp) RegisterUpgradeHandlers() {
	planName := "integration-test-upgrade"
	app.UpgradeKeeper.SetUpgradeHandler(planName, func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		return app.mm.RunMigrations(ctx, app.configurator, fromVM)
	})
}
