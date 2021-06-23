<template>
  <div class="block" :class="$style.header">
    <button
      :class="selected('/') ? $style.selected : ''"
      class="clickable"
      @click="$router.push('/')"
    >
      Main
    </button>
    <button
      :class="selected('/db') ? $style.selected : ''"
      class="clickable"
      @click="$router.push('/db')"
    >
      DB
    </button>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { RestApi } from '../util/RestApi';

export default defineComponent({
  components: {},
  async mounted() {},
  methods: {
    selected(path: string) {
      if (this.$route.path === path) return true;
      else return false;
    },
  },
  data: () => {
    return {};
  },
});
</script>

<style module lang="scss">
.header {
  display: flex;

  button {
    background: #2e2e2e;
    border-radius: 4px;
    border: 0;
    padding: 5px 15px;
    margin-right: 10px;
    color: #a5a5a5;

    &.selected {
      background: #d4670b;
      color: #fefefe;
    }
  }
}
</style>
