<template>
  <div class="block">
    <div class="header">
      Fap
      <img @click="add()" class="clickable" src="../../asset/add.svg" alt="" />
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
import { RestApi } from '../../util/RestApi';

export default defineComponent({
  props: {
    date: String,
  },
  watch: {
    date(value: Date) {
      this.refresh();
    },
  },
  async mounted() {
    this.refresh();
  },
  methods: {
    async refresh() {
      this.list = await RestApi.fap.getList();
    },
    async add() {
      const amount = prompt('Amount?');
      if (amount) {
        await RestApi.fap.add(Number(amount), Moment(this.date).format('YYYY-MM-DD'));
        this.refresh();
      }
    },
    currentWeight() {
      for (let i = 0; i < this.list.length; i++) {
        if (
          Moment(this.list[i].created).format('YYYY-MM-DD') ===
          Moment(this.date).format('YYYY-MM-DD')
        ) {
          return this.list[i].amount;
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
