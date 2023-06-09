package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"university_chain_de/testutil/sample"
)

func TestMsgStartErasmus_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgStartErasmus
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgStartErasmus{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgStartErasmus{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
