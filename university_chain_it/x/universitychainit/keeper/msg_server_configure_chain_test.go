package keeper_test

import (
	"context"
	"log"
	"os"
	"testing"
	keepertest "university_chain_it/testutil/keeper"
	"university_chain_it/x/universitychainit"
	"university_chain_it/x/universitychainit/keeper"
	"university_chain_it/x/universitychainit/testutil"
	"university_chain_it/x/universitychainit/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func setupMsgServerConfigureChain(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context, *gomock.Controller, *testutil.MockBankKeeperComplete) {
	ctrl := gomock.NewController(t)
	bankMock := testutil.NewMockBankKeeperComplete(ctrl)
	k, ctx := keepertest.UniversitychainitWithMocks(t, bankMock)
	universitychainit.InitGenesis(ctx, *k, *types.DefaultGenesis())
	server := keeper.NewMsgServerImpl(*k)
	context := sdk.WrapSDKContext(ctx)
	return server, *k, context, ctrl, bankMock
}

func TestConfigureChain(t *testing.T) {
	msgServer, _, context, _, _ := setupMsgServerConfigureChain(t)
	err := os.Chdir("/university_chain_it")
	if err != nil {
		log.Println(err)
	}

	createResponse, err := msgServer.ConfigureChain(context, &types.MsgConfigureChain{
		Creator: testutil.Chain_admin,
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgConfigureChainResponse{
		Status: 0,
	}, *createResponse)
}

func TestConfigureChainWrongCreator(t *testing.T) {
	msgServer, _, context, _, _ := setupMsgServerConfigureChain(t)
	createResponse, err := msgServer.ConfigureChain(context, &types.MsgConfigureChain{
		Creator: testutil.Mario_Rossi,
	})
	require.EqualError(t,
		err,
		"the key entered does not match the chain administrator's key")
	require.EqualValues(t, types.MsgConfigureChainResponse{
		Status: -1,
	}, *createResponse)
}

func TestConfigureChainJustConfigured(t *testing.T) {
	msgServer, _, context, _, _ := setupMsgServerConfigureChain(t)

	createResponse, err := msgServer.ConfigureChain(context, &types.MsgConfigureChain{
		Creator: testutil.Chain_admin,
	})
	require.Nil(t, err)

	createResponse, err = msgServer.ConfigureChain(context, &types.MsgConfigureChain{
		Creator: testutil.Chain_admin,
	})

	require.EqualError(t,
		err,
		"the initial configuration of the chain has already been performed")
	require.EqualValues(t, types.MsgConfigureChainResponse{
		Status: -1,
	}, *createResponse)
}
