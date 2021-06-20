<template>
  <router-view />
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
