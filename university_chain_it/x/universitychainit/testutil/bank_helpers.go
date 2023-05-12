package testutil

import (
	"context"
	"university_chain_it/x/universitychainit/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
)

func (bank *MockBankKeeperComplete) ExpectAny(context context.Context) {
	bank.EXPECT().SendCoinsFromAccountToModule(sdk.UnwrapSDKContext(context), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
	bank.EXPECT().SendCoinsFromModuleToAccount(sdk.UnwrapSDKContext(context), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
}

func coinsOf(amount uint64) sdk.Coins {
	return sdk.Coins{
		sdk.Coin{
			Denom:  sdk.DefaultBondDenom,
			Amount: sdk.NewInt(int64(amount)),
		},
	}
}

func (bank *MockBankKeeperComplete) PayTaxes(context context.Context, from string, to string, amount uint64) {
	fromAddr, err := sdk.AccAddressFromBech32(from)
	if err != nil {
		panic(err)
	}

	toAddr, err := sdk.AccAddressFromBech32(to)
	if err != nil {
		panic(err)
	}

	bank.EXPECT().SendCoinsFromAccountToModule(sdk.UnwrapSDKContext(context), fromAddr, types.ModuleName, coinsOf(amount)).Return(nil).AnyTimes()
	bank.EXPECT().SendCoinsFromModuleToAccount(sdk.UnwrapSDKContext(context), types.ModuleName, toAddr, coinsOf(amount)).Return(nil).AnyTimes()
}
