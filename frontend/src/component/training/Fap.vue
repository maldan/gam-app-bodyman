<template>
  <div class="block">
    <div class="header">
      Fap
      <img
        @click="
          $store.dispatch('modal/show', {
            name: 'prompt',
            data: { title: 'Add...', value: '', date: date },
            func: () => {
              $store.dispatch('fap/add');
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
    this.$store.dispatch('fap/getList');
  },
  methods: {
    currentWeight() {
      for (let i = 0; i < this.$store.state.fap.list.length; i++) {
        if (
          Moment(this.$store.state.fap.list[i].created).format('YYYY-MM-DD') ===
          Moment(this.date).format('YYYY-MM-DD')
        ) {
          return this.$store.state.fap.list[i].amount;
        }
      }
    },
  },
  data: () => {
    return {
      list: [] as any[],
    };
  },
});
</script>

<style lang="scss" module>
.weight {
}
</style>
