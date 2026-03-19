package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"university_chain_it/testutil/sample"
)

func TestMsgInsertStudentContactInfo_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgInsertStudentContactInfo
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgInsertStudentContactInfo{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgInsertStudentContactInfo{
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
