<template>
  <div class="main">
    <Header />

    <!-- Eat History -->
    <div style="display: flex; height: calc(100% - 60px); margin-top: 10px">
      <ExerciseList />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import Header from '../component/Header.vue';
import Input from '../component/Input.vue';
import ExerciseList from '../component/exercise/List.vue';

export default defineComponent({
  components: { Header, Input, ExerciseList },
  async mounted() {},
  methods: {},
  data: () => {
    return {};
  },
});
</script>

<style lang="scss" scoped>
.main {
  box-sizing: border-box;
  height: 100%;
  padding: 10px;
}
</style>
