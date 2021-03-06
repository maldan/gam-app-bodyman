import Axios from 'axios';

const API_URL = process.env.VUE_APP_API_URL || `${window.location.origin}/api`;

export const RestApi = {
  exercise: {
    async getList() {
      return (await Axios.get(`${API_URL}/exercise/list`)).data.response;
    },
    async get(id: string) {
      return (await Axios.get(`${API_URL}/exercise?id=${id}`)).data.response;
    },
    async findByName(name: string) {
      return (await Axios.get(`${API_URL}/exercise/byName?name=${name}`)).data.response;
    },
    async add(name: string, tool: string, muscleList: string[]) {
      return (
        await Axios.post(`${API_URL}/exercise`, {
          name,
          tool,
          muscleList,
        })
      ).data.response;
    },
    async update(id: string, name: string, tool: string, muscleList: string[]) {
      return (
        await Axios.patch(`${API_URL}/exercise`, {
          id,
          name,
          tool,
          muscleList,
        })
      ).data.response;
    },
    async delete(id: string) {
      return (await Axios.delete(`${API_URL}/exercise?id=${id}`)).data.response;
    },
  },
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
    async add(name: string, protein: string, carbohydrate: string, fat: string) {
      return (
        await Axios.post(`${API_URL}/product`, {
          name,
          protein,
          carbohydrate,
          fat,
        })
      ).data.response;
    },
    async update(id: string, name: string, protein: string, carbohydrate: string, fat: string) {
      return (
        await Axios.patch(`${API_URL}/product`, {
          id,
          name,
          protein,
          carbohydrate,
          fat,
        })
      ).data.response;
    },
    async delete(id: string) {
      return (await Axios.delete(`${API_URL}/product?id=${id}`)).data.response;
    },
  },
  eat: {
    async getYearMap(date: string) {
      return (await Axios.get(`${API_URL}/eat/yearMap?date=${date}`)).data.response;
    },
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
    async delete(id: string) {
      return (await Axios.delete(`${API_URL}/eat?id=${id}`)).data.response;
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
    async add(name: string, type: string) {
      return (
        await Axios.post(`${API_URL}/component`, {
          name,
          type,
        })
      ).data.response;
    },
    async getList() {
      return (await Axios.get(`${API_URL}/component/list`)).data.response;
    },
    async get(id: string) {
      return (await Axios.get(`${API_URL}/component?id=${id}`)).data.response;
    },
    async update(id: string, name: string, type: string) {
      return (
        await Axios.patch(`${API_URL}/component`, {
          id,
          name,
          type,
        })
      ).data.response;
    },
    async delete(id: string) {
      return (await Axios.delete(`${API_URL}/component?id=${id}`)).data.response;
    },
  },
  note: {
    async getList() {
      return (await Axios.get(`${API_URL}/note/list`)).data.response;
    },
    async get(id: string) {
      return (await Axios.get(`${API_URL}/note?id=${id}`)).data.response;
    },
    async update(note: any) {
      return (await Axios.patch(`${API_URL}/note`, note)).data.response;
    },
    async add(description: string) {
      return (
        await Axios.post(`${API_URL}/note`, {
          description,
        })
      ).data.response;
    },
    async delete(id: string) {
      return (await Axios.delete(`${API_URL}/note?id=${id}`)).data.response;
    },
  },
  training: {
    async getList() {
      return (await Axios.get(`${API_URL}/training/list`)).data.response;
    },
    async get(id: string) {
      return (await Axios.get(`${API_URL}/training?id=${id}`)).data.response;
    },
    async getFilterByDate(date: string) {
      return (await Axios.get(`${API_URL}/training/filterByDate?date=${date}`)).data.response;
    },
    async getTotalStatByDate(date: string) {
      return (await Axios.get(`${API_URL}/training/totalStatByDate?date=${date}`)).data.response;
    },
    async getYearMap(date: string) {
      return (await Axios.get(`${API_URL}/training/yearMap?date=${date}`)).data.response;
    },
    async update(
      id: string,
      exerciseId: string,
      reps: number,
      weight: number,
      distance: number,
      duration: number,
      created: string,
    ) {
      return (
        await Axios.patch(`${API_URL}/training`, {
          id,
          exerciseId,
          reps,
          weight,
          distance,
          duration,
          created,
        })
      ).data.response;
    },
    async add(
      exerciseId: string,
      reps: number,
      weight: number,
      distance: number,
      duration: number,
      created: string,
    ) {
      return (
        await Axios.post(`${API_URL}/training`, {
          exerciseId,
          reps,
          weight,
          distance,
          duration,
          created,
        })
      ).data.response;
    },
    async delete(id: string) {
      return (await Axios.delete(`${API_URL}/training?id=${id}`)).data.response;
    },
  },
};
