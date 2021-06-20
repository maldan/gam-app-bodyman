import Axios from 'axios';

const API_URL = process.env.VUE_APP_API_URL || `${window.location.origin}/api`;

export const RestApi = {
  product: {
    async getList() {
      return (await Axios.get(`${API_URL}/product/list`)).data.response;
    },
    async get(id: string) {
      return (await Axios.get(`${API_URL}/product?id=${id}`)).data.response;
    },
    async findByName(name: string) {
      return (await Axios.get(`${API_URL}/product/byName?name=${name}`)).data.response;
    },
  },
  eat: {
    async getTotalStatByDate(date: string) {
      return (await Axios.get(`${API_URL}/eat/totalStatByDate?date=${date}`)).data.response;
    },
    async getFilterByDate(date: string) {
      return (await Axios.get(`${API_URL}/eat/filterByDate?date=${date}`)).data.response;
    },
    async getList() {
      return (await Axios.get(`${API_URL}/eat/list`)).data.response;
    },
    async get(id: string) {
      return (await Axios.get(`${API_URL}/eat?id=${id}`)).data.response;
    },
    async add(productId: string, amount: string, created: string) {
      return (
        await Axios.post(`${API_URL}/eat`, {
          productId,
          amount,
          created,
        })
      ).data.response;
    },
    async update(id: string, productId: string, amount: string, created: string) {
      return (
        await Axios.patch(`${API_URL}/eat`, {
          id,
          productId,
          amount,
          created,
        })
      ).data.response;
    },
  },
  component: {
    async getList() {
      return (await Axios.get(`${API_URL}/component/list`)).data.response;
    },
    async get(id: string) {
      return (await Axios.get(`${API_URL}/component?id=${id}`)).data.response;
    },
  },
};
