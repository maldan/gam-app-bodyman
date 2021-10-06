<template>
  <router-view />

  <ui-modal v-if="$store.state.modal.name">
    <component :is="'modal-' + $store.state.modal.name" />
  </ui-modal>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import Moment from 'moment';

export default defineComponent({
  components: {},
  async mounted() {},
  methods: {},
  data: () => {
    return {
      moment: Moment,
    };
  },
});
</script>

<style lang="scss" scoped></style>
