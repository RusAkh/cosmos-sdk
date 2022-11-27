package runtime

import (
	runtimev1alpha1 "cosmossdk.io/api/cosmos/app/runtime/v1alpha1"
	"cosmossdk.io/core/appmodule"
)

func init() {
	appmodule.Register(&runtimev1alpha1.Module{})
}
