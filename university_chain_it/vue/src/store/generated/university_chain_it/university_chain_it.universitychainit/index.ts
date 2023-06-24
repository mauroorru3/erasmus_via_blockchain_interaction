import { txClient, queryClient, MissingWalletError , registry} from './module'

import { ChainInfo } from "./module/types/universitychainit/chain_info"
import { ContactInfo } from "./module/types/universitychainit/contact_info"
import { ErasmusInfo } from "./module/types/universitychainit/erasmus_info"
import { ForeignUniversities } from "./module/types/universitychainit/foreign_universities"
import { UniversitychainitPacketData } from "./module/types/universitychainit/packet"
import { NoData } from "./module/types/universitychainit/packet"
import { ErasmusStudentPacketData } from "./module/types/universitychainit/packet"
import { ErasmusStudentPacketAck } from "./module/types/universitychainit/packet"
import { ErasmusIndexPacketData } from "./module/types/universitychainit/packet"
import { ErasmusIndexPacketAck } from "./module/types/universitychainit/packet"
import { EndErasmusPeriodRequestPacketData } from "./module/types/universitychainit/packet"
import { EndErasmusPeriodRequestPacketAck } from "./module/types/universitychainit/packet"
import { FinalErasmusDataPacketData } from "./module/types/universitychainit/packet"
import { FinalErasmusDataPacketAck } from "./module/types/universitychainit/packet"
import { Params } from "./module/types/universitychainit/params"
import { PersonalInfo } from "./module/types/universitychainit/personal_info"
import { ProfessorsExams } from "./module/types/universitychainit/professors_exams"
import { ResidenceInfo } from "./module/types/universitychainit/residence_info"
import { StoredStudent } from "./module/types/universitychainit/stored_student"
import { StudentInfo } from "./module/types/universitychainit/student_info"
import { TaxesInfo } from "./module/types/universitychainit/taxes_info"
import { TranscriptOfRecords } from "./module/types/universitychainit/transcript_of_records"
import { UniversityInfo } from "./module/types/universitychainit/university_info"


export { ChainInfo, ContactInfo, ErasmusInfo, ForeignUniversities, UniversitychainitPacketData, NoData, ErasmusStudentPacketData, ErasmusStudentPacketAck, ErasmusIndexPacketData, ErasmusIndexPacketAck, EndErasmusPeriodRequestPacketData, EndErasmusPeriodRequestPacketAck, FinalErasmusDataPacketData, FinalErasmusDataPacketAck, Params, PersonalInfo, ProfessorsExams, ResidenceInfo, StoredStudent, StudentInfo, TaxesInfo, TranscriptOfRecords, UniversityInfo };

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
				ChainInfo: {},
				ForeignUniversities: {},
				ForeignUniversitiesAll: {},
				UniversityInfo: {},
				UniversityInfoAll: {},
				ProfessorsExams: {},
				ProfessorsExamsAll: {},
				StoredStudent: {},
				StoredStudentAll: {},
				
				_Structure: {
						ChainInfo: getStructure(ChainInfo.fromPartial({})),
						ContactInfo: getStructure(ContactInfo.fromPartial({})),
						ErasmusInfo: getStructure(ErasmusInfo.fromPartial({})),
						ForeignUniversities: getStructure(ForeignUniversities.fromPartial({})),
						UniversitychainitPacketData: getStructure(UniversitychainitPacketData.fromPartial({})),
						NoData: getStructure(NoData.fromPartial({})),
						ErasmusStudentPacketData: getStructure(ErasmusStudentPacketData.fromPartial({})),
						ErasmusStudentPacketAck: getStructure(ErasmusStudentPacketAck.fromPartial({})),
						ErasmusIndexPacketData: getStructure(ErasmusIndexPacketData.fromPartial({})),
						ErasmusIndexPacketAck: getStructure(ErasmusIndexPacketAck.fromPartial({})),
						EndErasmusPeriodRequestPacketData: getStructure(EndErasmusPeriodRequestPacketData.fromPartial({})),
						EndErasmusPeriodRequestPacketAck: getStructure(EndErasmusPeriodRequestPacketAck.fromPartial({})),
						FinalErasmusDataPacketData: getStructure(FinalErasmusDataPacketData.fromPartial({})),
						FinalErasmusDataPacketAck: getStructure(FinalErasmusDataPacketAck.fromPartial({})),
						Params: getStructure(Params.fromPartial({})),
						PersonalInfo: getStructure(PersonalInfo.fromPartial({})),
						ProfessorsExams: getStructure(ProfessorsExams.fromPartial({})),
						ResidenceInfo: getStructure(ResidenceInfo.fromPartial({})),
						StoredStudent: getStructure(StoredStudent.fromPartial({})),
						StudentInfo: getStructure(StudentInfo.fromPartial({})),
						TaxesInfo: getStructure(TaxesInfo.fromPartial({})),
						TranscriptOfRecords: getStructure(TranscriptOfRecords.fromPartial({})),
						UniversityInfo: getStructure(UniversityInfo.fromPartial({})),
						
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
				getChainInfo: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ChainInfo[JSON.stringify(params)] ?? {}
		},
				getForeignUniversities: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ForeignUniversities[JSON.stringify(params)] ?? {}
		},
				getForeignUniversitiesAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ForeignUniversitiesAll[JSON.stringify(params)] ?? {}
		},
				getUniversityInfo: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.UniversityInfo[JSON.stringify(params)] ?? {}
		},
				getUniversityInfoAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.UniversityInfoAll[JSON.stringify(params)] ?? {}
		},
				getProfessorsExams: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ProfessorsExams[JSON.stringify(params)] ?? {}
		},
				getProfessorsExamsAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ProfessorsExamsAll[JSON.stringify(params)] ?? {}
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
				
		getTypeStructure: (state) => (type) => {
			return state._Structure[type].fields
		},
		getRegistry: (state) => {
			return state._Registry
		}
	},
	actions: {
		init({ dispatch, rootGetters }) {
			console.log('Vuex module: university_chain_it.universitychainit initialized!')
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
		
		
		
		
		 		
		
		
		async QueryForeignUniversities({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryForeignUniversities( key.universityName)).data
				
					
				commit('QUERY', { query: 'ForeignUniversities', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryForeignUniversities', payload: { options: { all }, params: {...key},query }})
				return getters['getForeignUniversities']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryForeignUniversities API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryForeignUniversitiesAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryForeignUniversitiesAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryForeignUniversitiesAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'ForeignUniversitiesAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryForeignUniversitiesAll', payload: { options: { all }, params: {...key},query }})
				return getters['getForeignUniversitiesAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryForeignUniversitiesAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryUniversityInfo({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryUniversityInfo( key.universityName)).data
				
					
				commit('QUERY', { query: 'UniversityInfo', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryUniversityInfo', payload: { options: { all }, params: {...key},query }})
				return getters['getUniversityInfo']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryUniversityInfo API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryUniversityInfoAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryUniversityInfoAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryUniversityInfoAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'UniversityInfoAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryUniversityInfoAll', payload: { options: { all }, params: {...key},query }})
				return getters['getUniversityInfoAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryUniversityInfoAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryProfessorsExams({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryProfessorsExams( key.examName)).data
				
					
				commit('QUERY', { query: 'ProfessorsExams', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryProfessorsExams', payload: { options: { all }, params: {...key},query }})
				return getters['getProfessorsExams']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryProfessorsExams API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryProfessorsExamsAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryProfessorsExamsAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryProfessorsExamsAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'ProfessorsExamsAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryProfessorsExamsAll', payload: { options: { all }, params: {...key},query }})
				return getters['getProfessorsExamsAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryProfessorsExamsAll API Node Unavailable. Could not perform query: ' + e.message)
				
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
		async sendMsgInsertExamGrade({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgInsertExamGrade(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgInsertExamGrade:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgInsertExamGrade:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgStartErasmus({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgStartErasmus(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgStartErasmus:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgStartErasmus:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgPayTaxes({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgPayTaxes(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgPayTaxes:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgPayTaxes:Send Could not broadcast Tx: '+ e.message)
				}
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
		async sendMsgInsertStudentContactInfo({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgInsertStudentContactInfo(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgInsertStudentContactInfo:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgInsertStudentContactInfo:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgInsertStudentPersonalInfo({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgInsertStudentPersonalInfo(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgInsertStudentPersonalInfo:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgInsertStudentPersonalInfo:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgSendEndErasmusPeriodRequest({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSendEndErasmusPeriodRequest(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSendEndErasmusPeriodRequest:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgSendEndErasmusPeriodRequest:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgInsertErasmusRequest({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgInsertErasmusRequest(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgInsertErasmusRequest:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgInsertErasmusRequest:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgInsertErasmusExam({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgInsertErasmusExam(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgInsertErasmusExam:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgInsertErasmusExam:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgRegisterNewStudent({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgRegisterNewStudent(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRegisterNewStudent:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgRegisterNewStudent:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgInsertStudentResidenceInfo({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgInsertStudentResidenceInfo(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgInsertStudentResidenceInfo:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgInsertStudentResidenceInfo:Send Could not broadcast Tx: '+ e.message)
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
		async MsgInsertExamGrade({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgInsertExamGrade(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgInsertExamGrade:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgInsertExamGrade:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgStartErasmus({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgStartErasmus(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgStartErasmus:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgStartErasmus:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgPayTaxes({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgPayTaxes(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgPayTaxes:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgPayTaxes:Create Could not create message: ' + e.message)
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
		async MsgInsertStudentContactInfo({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgInsertStudentContactInfo(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgInsertStudentContactInfo:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgInsertStudentContactInfo:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgInsertStudentPersonalInfo({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgInsertStudentPersonalInfo(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgInsertStudentPersonalInfo:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgInsertStudentPersonalInfo:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgSendEndErasmusPeriodRequest({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSendEndErasmusPeriodRequest(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSendEndErasmusPeriodRequest:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgSendEndErasmusPeriodRequest:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgInsertErasmusRequest({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgInsertErasmusRequest(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgInsertErasmusRequest:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgInsertErasmusRequest:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgInsertErasmusExam({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgInsertErasmusExam(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgInsertErasmusExam:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgInsertErasmusExam:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgRegisterNewStudent({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgRegisterNewStudent(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRegisterNewStudent:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgRegisterNewStudent:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgInsertStudentResidenceInfo({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgInsertStudentResidenceInfo(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgInsertStudentResidenceInfo:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgInsertStudentResidenceInfo:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
