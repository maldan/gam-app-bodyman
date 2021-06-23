<template>
  <div class="main">
    <Header />

    <!-- Eat History -->
    <div style="display: flex; height: calc(100% - 60px); margin-top: 10px">
      <!-- Food list -->
      <ProductList style="margin-right: 10px" />
      <ComponentList />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import Header from '../component/Header.vue';
import Input from '../component/Input.vue';
import ProductList from '../component/product/ProductList.vue';
import ComponentList from '../component/component/List.vue';

export default defineComponent({
  components: { Header, Input, ProductList, ComponentList },
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
