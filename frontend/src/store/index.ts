import { createStore } from 'vuex';
import main, { MainStore } from './main';
import weight, { WeightStore } from './weight';
import fap, { FapStore } from './fap';
import modal, { ModalStore } from '../gam_sdk_ui/vue/store/modal';

export type MainTree = {
  main: MainStore;
  modal: ModalStore;
  weight: WeightStore;
  fap: FapStore;
};

export default createStore({
  modules: { main, weight, fap, modal },
});
