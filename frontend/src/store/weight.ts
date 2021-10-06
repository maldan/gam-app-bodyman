import { ActionContext } from 'vuex';
import { MainTree } from '.';
import Axios from 'axios';
import Moment from 'moment';

export type WeightStore = {
  list: [];
};
export type WeightActionContext = ActionContext<WeightStore, MainTree>;

export default {
  namespaced: true,
  state: {
    list: [],
  },
  mutations: {
    SET_LIST(state: WeightStore, payload: any) {
      state.list = payload;
    },
  },
  actions: {
    async getList(action: WeightActionContext) {
      const list = (await Axios.get(`${action.rootState.main.API_URL}/weight/list`)).data.response;
      action.commit('SET_LIST', list);
    },
    async add(action: WeightActionContext) {
      await Axios.post(`${action.rootState.main.API_URL}/weight`, {
        weight: action.rootState.modal.data.value * 1,
        created: Moment().format('YYYY-MM-DD'),
      });
      action.dispatch('getList');
    },
  },
};
