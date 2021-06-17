<template>
  <div class="main"></div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { RestApi } from '../util/RestApi';

export default defineComponent({
  components: {},
  async mounted() {
    console.log(await RestApi.food.getList());
  },
  methods: {},
  data: () => {
    return {};
  },
});
</script>

<style lang="scss" scoped>
.main {
}
</style>
