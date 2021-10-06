export type MainStore = {
  API_URL: string;
};

export default {
  namespaced: true,
  state: {
    API_URL: process.env.VUE_APP_API_URL || `${window.location.origin}/api`,
  },
  mutations: {},
  actions: {
    getList() {},
  },
};
