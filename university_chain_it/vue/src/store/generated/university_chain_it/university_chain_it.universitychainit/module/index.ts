// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgInsertStudentContactInfo } from "./types/universitychainit/tx";
import { MsgStartErasmus } from "./types/universitychainit/tx";
import { MsgRegisterNewStudent } from "./types/universitychainit/tx";
import { MsgConfigureChain } from "./types/universitychainit/tx";
import { MsgInsertErasmusRequest } from "./types/universitychainit/tx";
import { MsgInsertExamGrade } from "./types/universitychainit/tx";
import { MsgSendEndErasmusPeriodRequest } from "./types/universitychainit/tx";
import { MsgPayTaxes } from "./types/universitychainit/tx";
import { MsgInsertErasmusExam } from "./types/universitychainit/tx";
import { MsgSendErasmusStudent } from "./types/universitychainit/tx";
import { MsgInsertStudentPersonalInfo } from "./types/universitychainit/tx";
import { MsgInsertStudentResidenceInfo } from "./types/universitychainit/tx";


const types = [
  ["/university_chain_it.universitychainit.MsgInsertStudentContactInfo", MsgInsertStudentContactInfo],
  ["/university_chain_it.universitychainit.MsgStartErasmus", MsgStartErasmus],
  ["/university_chain_it.universitychainit.MsgRegisterNewStudent", MsgRegisterNewStudent],
  ["/university_chain_it.universitychainit.MsgConfigureChain", MsgConfigureChain],
  ["/university_chain_it.universitychainit.MsgInsertErasmusRequest", MsgInsertErasmusRequest],
  ["/university_chain_it.universitychainit.MsgInsertExamGrade", MsgInsertExamGrade],
  ["/university_chain_it.universitychainit.MsgSendEndErasmusPeriodRequest", MsgSendEndErasmusPeriodRequest],
  ["/university_chain_it.universitychainit.MsgPayTaxes", MsgPayTaxes],
  ["/university_chain_it.universitychainit.MsgInsertErasmusExam", MsgInsertErasmusExam],
  ["/university_chain_it.universitychainit.MsgSendErasmusStudent", MsgSendErasmusStudent],
  ["/university_chain_it.universitychainit.MsgInsertStudentPersonalInfo", MsgInsertStudentPersonalInfo],
  ["/university_chain_it.universitychainit.MsgInsertStudentResidenceInfo", MsgInsertStudentResidenceInfo],
  
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
    msgInsertStudentContactInfo: (data: MsgInsertStudentContactInfo): EncodeObject => ({ typeUrl: "/university_chain_it.universitychainit.MsgInsertStudentContactInfo", value: MsgInsertStudentContactInfo.fromPartial( data ) }),
    msgStartErasmus: (data: MsgStartErasmus): EncodeObject => ({ typeUrl: "/university_chain_it.universitychainit.MsgStartErasmus", value: MsgStartErasmus.fromPartial( data ) }),
    msgRegisterNewStudent: (data: MsgRegisterNewStudent): EncodeObject => ({ typeUrl: "/university_chain_it.universitychainit.MsgRegisterNewStudent", value: MsgRegisterNewStudent.fromPartial( data ) }),
    msgConfigureChain: (data: MsgConfigureChain): EncodeObject => ({ typeUrl: "/university_chain_it.universitychainit.MsgConfigureChain", value: MsgConfigureChain.fromPartial( data ) }),
    msgInsertErasmusRequest: (data: MsgInsertErasmusRequest): EncodeObject => ({ typeUrl: "/university_chain_it.universitychainit.MsgInsertErasmusRequest", value: MsgInsertErasmusRequest.fromPartial( data ) }),
    msgInsertExamGrade: (data: MsgInsertExamGrade): EncodeObject => ({ typeUrl: "/university_chain_it.universitychainit.MsgInsertExamGrade", value: MsgInsertExamGrade.fromPartial( data ) }),
    msgSendEndErasmusPeriodRequest: (data: MsgSendEndErasmusPeriodRequest): EncodeObject => ({ typeUrl: "/university_chain_it.universitychainit.MsgSendEndErasmusPeriodRequest", value: MsgSendEndErasmusPeriodRequest.fromPartial( data ) }),
    msgPayTaxes: (data: MsgPayTaxes): EncodeObject => ({ typeUrl: "/university_chain_it.universitychainit.MsgPayTaxes", value: MsgPayTaxes.fromPartial( data ) }),
    msgInsertErasmusExam: (data: MsgInsertErasmusExam): EncodeObject => ({ typeUrl: "/university_chain_it.universitychainit.MsgInsertErasmusExam", value: MsgInsertErasmusExam.fromPartial( data ) }),
    msgSendErasmusStudent: (data: MsgSendErasmusStudent): EncodeObject => ({ typeUrl: "/university_chain_it.universitychainit.MsgSendErasmusStudent", value: MsgSendErasmusStudent.fromPartial( data ) }),
    msgInsertStudentPersonalInfo: (data: MsgInsertStudentPersonalInfo): EncodeObject => ({ typeUrl: "/university_chain_it.universitychainit.MsgInsertStudentPersonalInfo", value: MsgInsertStudentPersonalInfo.fromPartial( data ) }),
    msgInsertStudentResidenceInfo: (data: MsgInsertStudentResidenceInfo): EncodeObject => ({ typeUrl: "/university_chain_it.universitychainit.MsgInsertStudentResidenceInfo", value: MsgInsertStudentResidenceInfo.fromPartial( data ) }),
    
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
