package cli_test

import (
	"fmt"
	"testing"

	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"google.golang.org/grpc/status"

	"hub/testutil/network"
	"hub/testutil/nullify"
	"hub/x/hub/client/cli"
	"hub/x/hub/types"
)

func networkWithTranscriptOfRecordsObjects(t *testing.T) (*network.Network, types.TranscriptOfRecords) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	transcriptOfRecords := &types.TranscriptOfRecords{}
	nullify.Fill(&transcriptOfRecords)
	state.TranscriptOfRecords = transcriptOfRecords
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), *state.TranscriptOfRecords
}

func TestShowTranscriptOfRecords(t *testing.T) {
	net, obj := networkWithTranscriptOfRecordsObjects(t)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc string
		args []string
		err  error
		obj  types.TranscriptOfRecords
	}{
		{
			desc: "get",
			args: common,
			obj:  obj,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			var args []string
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowTranscriptOfRecords(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetTranscriptOfRecordsResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.TranscriptOfRecords)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.TranscriptOfRecords),
				)
			}
		})
	}
}
