import { txClient, queryClient, MissingWalletError , registry} from './module'

import { ChainInfo } from "./module/types/hub/chain_info"
import { ContactInfo } from "./module/types/hub/contact_info"
import { ErasmusInfo } from "./module/types/hub/erasmus_info"
import { HubPacketData } from "./module/types/hub/packet"
import { NoData } from "./module/types/hub/packet"
import { ErasmusStudentPacketData } from "./module/types/hub/packet"
import { ErasmusStudentPacketAck } from "./module/types/hub/packet"
import { Params } from "./module/types/hub/params"
import { PersonalInfo } from "./module/types/hub/personal_info"
import { ResidenceInfo } from "./module/types/hub/residence_info"
import { StoredStudent } from "./module/types/hub/stored_student"
import { StudentInfo } from "./module/types/hub/student_info"
import { TaxesInfo } from "./module/types/hub/taxes_info"
import { TranscriptOfRecords } from "./module/types/hub/transcript_of_records"
import { Universities } from "./module/types/hub/universities"


export { ChainInfo, ContactInfo, ErasmusInfo, HubPacketData, NoData, ErasmusStudentPacketData, ErasmusStudentPacketAck, Params, PersonalInfo, ResidenceInfo, StoredStudent, StudentInfo, TaxesInfo, TranscriptOfRecords, Universities };

async function initTxClient(vuexGetters) {
	return await txClient(vuexGetters['common/wallet/signer'], {
		addr: vuexGetters['common/env/apiTendermint']
	})
}

async function initQueryClient(vuexGetters) {
	return await queryClient({
		addr: vuexGetters['common/env/apiCosmos']
	})
}

function mergeResults(value, next_values) {
	for (let prop of Object.keys(next_values)) {
		if (Array.isArray(next_values[prop])) {
			value[prop]=[...value[prop], ...next_values[prop]]
		}else{
			value[prop]=next_values[prop]
		}
	}
	return value
}

function getStructure(template) {
	let structure = { fields: [] }
	for (const [key, value] of Object.entries(template)) {
		let field: any = {}
		field.name = key
		field.type = typeof value
		structure.fields.push(field)
	}
	return structure
}

const getDefaultState = () => {
	return {
				Params: {},
				StudentInfo: {},
				TranscriptOfRecords: {},
				PersonalInfo: {},
				ResidenceInfo: {},
				ContactInfo: {},
				TaxesInfo: {},
				ErasmusInfo: {},
				StoredStudent: {},
				StoredStudentAll: {},
				ChainInfo: {},
				Universities: {},
				UniversitiesAll: {},
				
				_Structure: {
						ChainInfo: getStructure(ChainInfo.fromPartial({})),
						ContactInfo: getStructure(ContactInfo.fromPartial({})),
						ErasmusInfo: getStructure(ErasmusInfo.fromPartial({})),
						HubPacketData: getStructure(HubPacketData.fromPartial({})),
						NoData: getStructure(NoData.fromPartial({})),
						ErasmusStudentPacketData: getStructure(ErasmusStudentPacketData.fromPartial({})),
						ErasmusStudentPacketAck: getStructure(ErasmusStudentPacketAck.fromPartial({})),
						Params: getStructure(Params.fromPartial({})),
						PersonalInfo: getStructure(PersonalInfo.fromPartial({})),
						ResidenceInfo: getStructure(ResidenceInfo.fromPartial({})),
						StoredStudent: getStructure(StoredStudent.fromPartial({})),
						StudentInfo: getStructure(StudentInfo.fromPartial({})),
						TaxesInfo: getStructure(TaxesInfo.fromPartial({})),
						TranscriptOfRecords: getStructure(TranscriptOfRecords.fromPartial({})),
						Universities: getStructure(Universities.fromPartial({})),
						
		},
		_Registry: registry,
		_Subscriptions: new Set(),
	}
}

// initial state
const state = getDefaultState()

export default {
	namespaced: true,
	state,
	mutations: {
		RESET_STATE(state) {
			Object.assign(state, getDefaultState())
		},
		QUERY(state, { query, key, value }) {
			state[query][JSON.stringify(key)] = value
		},
		SUBSCRIBE(state, subscription) {
			state._Subscriptions.add(JSON.stringify(subscription))
		},
		UNSUBSCRIBE(state, subscription) {
			state._Subscriptions.delete(JSON.stringify(subscription))
		}
	},
	getters: {
				getParams: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Params[JSON.stringify(params)] ?? {}
		},
				getStudentInfo: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.StudentInfo[JSON.stringify(params)] ?? {}
		},
				getTranscriptOfRecords: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.TranscriptOfRecords[JSON.stringify(params)] ?? {}
		},
				getPersonalInfo: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.PersonalInfo[JSON.stringify(params)] ?? {}
		},
				getResidenceInfo: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ResidenceInfo[JSON.stringify(params)] ?? {}
		},
				getContactInfo: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ContactInfo[JSON.stringify(params)] ?? {}
		},
				getTaxesInfo: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.TaxesInfo[JSON.stringify(params)] ?? {}
		},
				getErasmusInfo: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ErasmusInfo[JSON.stringify(params)] ?? {}
		},
				getStoredStudent: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.StoredStudent[JSON.stringify(params)] ?? {}
		},
				getStoredStudentAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.StoredStudentAll[JSON.stringify(params)] ?? {}
		},
				getChainInfo: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ChainInfo[JSON.stringify(params)] ?? {}
		},
				getUniversities: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Universities[JSON.stringify(params)] ?? {}
		},
				getUniversitiesAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.UniversitiesAll[JSON.stringify(params)] ?? {}
		},
				
		getTypeStructure: (state) => (type) => {
			return state._Structure[type].fields
		},
		getRegistry: (state) => {
			return state._Registry
		}
	},
	actions: {
		init({ dispatch, rootGetters }) {
			console.log('Vuex module: hub.hub initialized!')
			if (rootGetters['common/env/client']) {
				rootGetters['common/env/client'].on('newblock', () => {
					dispatch('StoreUpdate')
				})
			}
		},
		resetState({ commit }) {
			commit('RESET_STATE')
		},
		unsubscribe({ commit }, subscription) {
			commit('UNSUBSCRIBE', subscription)
		},
		async StoreUpdate({ state, dispatch }) {
			state._Subscriptions.forEach(async (subscription) => {
				try {
					const sub=JSON.parse(subscription)
					await dispatch(sub.action, sub.payload)
				}catch(e) {
					throw new Error('Subscriptions: ' + e.message)
				}
			})
		},
		
		
		
		 		
		
		
		async QueryParams({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryParams()).data
				
					
				commit('QUERY', { query: 'Params', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryParams', payload: { options: { all }, params: {...key},query }})
				return getters['getParams']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryParams API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryStudentInfo({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryStudentInfo()).data
				
					
				commit('QUERY', { query: 'StudentInfo', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryStudentInfo', payload: { options: { all }, params: {...key},query }})
				return getters['getStudentInfo']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryStudentInfo API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryTranscriptOfRecords({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryTranscriptOfRecords()).data
				
					
				commit('QUERY', { query: 'TranscriptOfRecords', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryTranscriptOfRecords', payload: { options: { all }, params: {...key},query }})
				return getters['getTranscriptOfRecords']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryTranscriptOfRecords API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryPersonalInfo({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryPersonalInfo()).data
				
					
				commit('QUERY', { query: 'PersonalInfo', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryPersonalInfo', payload: { options: { all }, params: {...key},query }})
				return getters['getPersonalInfo']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryPersonalInfo API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryResidenceInfo({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryResidenceInfo()).data
				
					
				commit('QUERY', { query: 'ResidenceInfo', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryResidenceInfo', payload: { options: { all }, params: {...key},query }})
				return getters['getResidenceInfo']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryResidenceInfo API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryContactInfo({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryContactInfo()).data
				
					
				commit('QUERY', { query: 'ContactInfo', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryContactInfo', payload: { options: { all }, params: {...key},query }})
				return getters['getContactInfo']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryContactInfo API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryTaxesInfo({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryTaxesInfo()).data
				
					
				commit('QUERY', { query: 'TaxesInfo', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryTaxesInfo', payload: { options: { all }, params: {...key},query }})
				return getters['getTaxesInfo']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryTaxesInfo API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryErasmusInfo({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryErasmusInfo()).data
				
					
				commit('QUERY', { query: 'ErasmusInfo', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryErasmusInfo', payload: { options: { all }, params: {...key},query }})
				return getters['getErasmusInfo']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryErasmusInfo API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryStoredStudent({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryStoredStudent( key.index)).data
				
					
				commit('QUERY', { query: 'StoredStudent', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryStoredStudent', payload: { options: { all }, params: {...key},query }})
				return getters['getStoredStudent']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryStoredStudent API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryStoredStudentAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryStoredStudentAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryStoredStudentAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'StoredStudentAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryStoredStudentAll', payload: { options: { all }, params: {...key},query }})
				return getters['getStoredStudentAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryStoredStudentAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryChainInfo({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryChainInfo()).data
				
					
				commit('QUERY', { query: 'ChainInfo', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryChainInfo', payload: { options: { all }, params: {...key},query }})
				return getters['getChainInfo']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryChainInfo API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryUniversities({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryUniversities( key.universityName)).data
				
					
				commit('QUERY', { query: 'Universities', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryUniversities', payload: { options: { all }, params: {...key},query }})
				return getters['getUniversities']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryUniversities API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryUniversitiesAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryUniversitiesAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryUniversitiesAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'UniversitiesAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryUniversitiesAll', payload: { options: { all }, params: {...key},query }})
				return getters['getUniversitiesAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryUniversitiesAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgConfigureChain({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgConfigureChain(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgConfigureChain:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgConfigureChain:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgSendErasmusStudent({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSendErasmusStudent(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSendErasmusStudent:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgSendErasmusStudent:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgConfigureChain({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgConfigureChain(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgConfigureChain:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgConfigureChain:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgSendErasmusStudent({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSendErasmusStudent(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSendErasmusStudent:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgSendErasmusStudent:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
