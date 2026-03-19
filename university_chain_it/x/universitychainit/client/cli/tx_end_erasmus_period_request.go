package cli

import (
	"strconv"

	"university_chain_it/x/universitychainit/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	channelutils "github.com/cosmos/ibc-go/v3/modules/core/04-channel/client/utils"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdSendEndErasmusPeriodRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send-end-erasmus-period-request [src-port] [src-channel] [starting-university-name] [destination-university-name] [index] [foreign-index]",
		Short: "Send a end_erasmus_period_request over IBC",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			creator := clientCtx.GetFromAddress().String()
			srcPort := args[0]
			srcChannel := args[1]

			argStartingUniversityName := args[2]
			argDestinationUniversityName := args[3]
			argIndex := args[4]
			argForeignIndex := args[5]

			// Get the relative timeout timestamp
			timeoutTimestamp, err := cmd.Flags().GetUint64(flagPacketTimeoutTimestamp)
			if err != nil {
				return err
			}
			consensusState, _, _, err := channelutils.QueryLatestConsensusState(clientCtx, srcPort, srcChannel)
			if err != nil {
				return err
			}
			if timeoutTimestamp != 0 {
				timeoutTimestamp = consensusState.GetTimestamp() + timeoutTimestamp
			}

			msg := types.NewMsgSendEndErasmusPeriodRequest(creator, srcPort, srcChannel, timeoutTimestamp, argStartingUniversityName, argDestinationUniversityName, argIndex, argForeignIndex)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().Uint64(flagPacketTimeoutTimestamp, DefaultRelativePacketTimeoutTimestamp, "Packet timeout timestamp in nanoseconds. Default is 10 minutes.")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
