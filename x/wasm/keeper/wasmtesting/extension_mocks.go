package wasmtesting

import (
	sdk "github.com/line/lbm-sdk/types"
	authtypes "github.com/line/lbm-sdk/x/auth/types"
)

type MockCoinTransferrer struct {
	TransferCoinsFn func(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
}

func (m *MockCoinTransferrer) AddToInactiveAddr(ctx sdk.Context, address sdk.AccAddress) {
	panic("implement me")
}

func (m *MockCoinTransferrer) DeleteFromInactiveAddr(ctx sdk.Context, address sdk.AccAddress) {
	panic("implement me")
}

func (m *MockCoinTransferrer) TransferCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error {
	if m.TransferCoinsFn == nil {
		panic("not expected to be called")
	}
	return m.TransferCoinsFn(ctx, fromAddr, toAddr, amt)
}

type AccountPrunerMock struct {
	CleanupExistingAccountFn func(ctx sdk.Context, existingAccount authtypes.AccountI) (handled bool, err error)
}

func (m AccountPrunerMock) CleanupExistingAccount(ctx sdk.Context, existingAccount authtypes.AccountI) (handled bool, err error) {
	if m.CleanupExistingAccountFn == nil {
		panic("not expected to be called")
	}
	return m.CleanupExistingAccountFn(ctx, existingAccount)
}
