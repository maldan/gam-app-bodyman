import Axios from 'axios';

const API_URL = process.env.VUE_APP_API_URL || `${window.location.origin}/api`;

export const RestApi = {
  food: {
    async getList() {
      return (await Axios.get(`${API_URL}/food/list`)).data.response;
    },
  },
};
