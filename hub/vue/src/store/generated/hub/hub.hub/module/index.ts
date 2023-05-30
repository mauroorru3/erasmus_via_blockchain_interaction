// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgSendErasmusStudent } from "./types/hub/tx";
import { MsgSendFinalErasmusData } from "./types/hub/tx";
import { MsgSendErasmusIndex } from "./types/hub/tx";
import { MsgConfigureChain } from "./types/hub/tx";
import { MsgSendEndErasmusPeriodRequest } from "./types/hub/tx";


const types = [
  ["/hub.hub.MsgSendErasmusStudent", MsgSendErasmusStudent],
  ["/hub.hub.MsgSendFinalErasmusData", MsgSendFinalErasmusData],
  ["/hub.hub.MsgSendErasmusIndex", MsgSendErasmusIndex],
  ["/hub.hub.MsgConfigureChain", MsgConfigureChain],
  ["/hub.hub.MsgSendEndErasmusPeriodRequest", MsgSendEndErasmusPeriodRequest],
  
];
export const MissingWalletError = new Error("wallet is required");

export const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;
  let client;
  if (addr) {
    client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  }else{
    client = await SigningStargateClient.offline( wallet, { registry });
  }
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgSendErasmusStudent: (data: MsgSendErasmusStudent): EncodeObject => ({ typeUrl: "/hub.hub.MsgSendErasmusStudent", value: MsgSendErasmusStudent.fromPartial( data ) }),
    msgSendFinalErasmusData: (data: MsgSendFinalErasmusData): EncodeObject => ({ typeUrl: "/hub.hub.MsgSendFinalErasmusData", value: MsgSendFinalErasmusData.fromPartial( data ) }),
    msgSendErasmusIndex: (data: MsgSendErasmusIndex): EncodeObject => ({ typeUrl: "/hub.hub.MsgSendErasmusIndex", value: MsgSendErasmusIndex.fromPartial( data ) }),
    msgConfigureChain: (data: MsgConfigureChain): EncodeObject => ({ typeUrl: "/hub.hub.MsgConfigureChain", value: MsgConfigureChain.fromPartial( data ) }),
    msgSendEndErasmusPeriodRequest: (data: MsgSendEndErasmusPeriodRequest): EncodeObject => ({ typeUrl: "/hub.hub.MsgSendEndErasmusPeriodRequest", value: MsgSendEndErasmusPeriodRequest.fromPartial( data ) }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
