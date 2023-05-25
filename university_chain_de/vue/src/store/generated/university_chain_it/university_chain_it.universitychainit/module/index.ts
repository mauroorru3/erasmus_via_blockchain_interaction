// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgInsertErasmusExam } from "./types/universitychainde/tx";
import { MsgInsertErasmusRequest } from "./types/universitychainde/tx";
import { MsgConfigureChain } from "./types/universitychainde/tx";
import { MsgInsertExamGrade } from "./types/universitychainde/tx";
import { MsgPayTaxes } from "./types/universitychainde/tx";
import { MsgSendEndErasmusPeriodRequest } from "./types/universitychainde/tx";
import { MsgSendErasmusStudent } from "./types/universitychainde/tx";
import { MsgStartErasmus } from "./types/universitychainde/tx";
import { MsgRegisterNewStudent } from "./types/universitychainde/tx";
import { MsgInsertStudentContactInfo } from "./types/universitychainde/tx";
import { MsgInsertStudentResidenceInfo } from "./types/universitychainde/tx";
import { MsgInsertStudentPersonalInfo } from "./types/universitychainde/tx";


const types = [
  ["/university_chain_de.universitychainde.MsgInsertErasmusExam", MsgInsertErasmusExam],
  ["/university_chain_de.universitychainde.MsgInsertErasmusRequest", MsgInsertErasmusRequest],
  ["/university_chain_de.universitychainde.MsgConfigureChain", MsgConfigureChain],
  ["/university_chain_de.universitychainde.MsgInsertExamGrade", MsgInsertExamGrade],
  ["/university_chain_de.universitychainde.MsgPayTaxes", MsgPayTaxes],
  ["/university_chain_de.universitychainde.MsgSendEndErasmusPeriodRequest", MsgSendEndErasmusPeriodRequest],
  ["/university_chain_de.universitychainde.MsgSendErasmusStudent", MsgSendErasmusStudent],
  ["/university_chain_de.universitychainde.MsgStartErasmus", MsgStartErasmus],
  ["/university_chain_de.universitychainde.MsgRegisterNewStudent", MsgRegisterNewStudent],
  ["/university_chain_de.universitychainde.MsgInsertStudentContactInfo", MsgInsertStudentContactInfo],
  ["/university_chain_de.universitychainde.MsgInsertStudentResidenceInfo", MsgInsertStudentResidenceInfo],
  ["/university_chain_de.universitychainde.MsgInsertStudentPersonalInfo", MsgInsertStudentPersonalInfo],
  
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
    msgInsertErasmusExam: (data: MsgInsertErasmusExam): EncodeObject => ({ typeUrl: "/university_chain_de.universitychainde.MsgInsertErasmusExam", value: MsgInsertErasmusExam.fromPartial( data ) }),
    msgInsertErasmusRequest: (data: MsgInsertErasmusRequest): EncodeObject => ({ typeUrl: "/university_chain_de.universitychainde.MsgInsertErasmusRequest", value: MsgInsertErasmusRequest.fromPartial( data ) }),
    msgConfigureChain: (data: MsgConfigureChain): EncodeObject => ({ typeUrl: "/university_chain_de.universitychainde.MsgConfigureChain", value: MsgConfigureChain.fromPartial( data ) }),
    msgInsertExamGrade: (data: MsgInsertExamGrade): EncodeObject => ({ typeUrl: "/university_chain_de.universitychainde.MsgInsertExamGrade", value: MsgInsertExamGrade.fromPartial( data ) }),
    msgPayTaxes: (data: MsgPayTaxes): EncodeObject => ({ typeUrl: "/university_chain_de.universitychainde.MsgPayTaxes", value: MsgPayTaxes.fromPartial( data ) }),
    msgSendEndErasmusPeriodRequest: (data: MsgSendEndErasmusPeriodRequest): EncodeObject => ({ typeUrl: "/university_chain_de.universitychainde.MsgSendEndErasmusPeriodRequest", value: MsgSendEndErasmusPeriodRequest.fromPartial( data ) }),
    msgSendErasmusStudent: (data: MsgSendErasmusStudent): EncodeObject => ({ typeUrl: "/university_chain_de.universitychainde.MsgSendErasmusStudent", value: MsgSendErasmusStudent.fromPartial( data ) }),
    msgStartErasmus: (data: MsgStartErasmus): EncodeObject => ({ typeUrl: "/university_chain_de.universitychainde.MsgStartErasmus", value: MsgStartErasmus.fromPartial( data ) }),
    msgRegisterNewStudent: (data: MsgRegisterNewStudent): EncodeObject => ({ typeUrl: "/university_chain_de.universitychainde.MsgRegisterNewStudent", value: MsgRegisterNewStudent.fromPartial( data ) }),
    msgInsertStudentContactInfo: (data: MsgInsertStudentContactInfo): EncodeObject => ({ typeUrl: "/university_chain_de.universitychainde.MsgInsertStudentContactInfo", value: MsgInsertStudentContactInfo.fromPartial( data ) }),
    msgInsertStudentResidenceInfo: (data: MsgInsertStudentResidenceInfo): EncodeObject => ({ typeUrl: "/university_chain_de.universitychainde.MsgInsertStudentResidenceInfo", value: MsgInsertStudentResidenceInfo.fromPartial( data ) }),
    msgInsertStudentPersonalInfo: (data: MsgInsertStudentPersonalInfo): EncodeObject => ({ typeUrl: "/university_chain_de.universitychainde.MsgInsertStudentPersonalInfo", value: MsgInsertStudentPersonalInfo.fromPartial( data ) }),
    
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
