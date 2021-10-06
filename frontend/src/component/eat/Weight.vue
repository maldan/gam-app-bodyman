<template>
  <div class="block">
    <div class="header">
      Weight
      <img
        @click="
          $store.dispatch('modal/show', {
            name: 'prompt',
            data: { title: 'New weight...', value: '' },
            func: () => {
              $store.dispatch('weight/add');
            },
          })
        "
        class="clickable"
        src="../../asset/add.svg"
        alt=""
      />
    </div>
    <div class="body">
      <div :class="$style.weight">
        {{ currentWeight() }}
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Moment from 'moment';
import { defineComponent } from 'vue';

export default defineComponent({
  props: {
    date: String,
  },
  async mounted() {
    this.$store.dispatch('weight/getList');
  },
  methods: {
    currentWeight() {
      for (let i = 0; i < this.$store.state.weight.list.length; i++) {
        if (
          Moment(this.$store.state.weight.list[i].created).format('YYYY-MM-DD') ===
          Moment(this.date).format('YYYY-MM-DD')
        ) {
          return this.$store.state.weight.list[i].weight;
        }
      }
    },
  },
  data: () => {
    return {};
  },
});
</script>

<style lang="scss" module>
.weight {
}
</style>
