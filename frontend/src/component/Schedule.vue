<template>
  <div :class="$style.schedule">
    <div :class="$style.year">
      <div :class="$style.month" v-for="(x, i) in month" :key="x">
        <div>{{ x }}</div>
        <div style="display: grid; margin-top: 10px; grid-template-columns: repeat(7, 15px)">
          <div
            @click="$emit('select', z), (date = z)"
            class="clickable"
            :class="[
              $style.cell,
              $root.moment(new Date()).format('DD MMM YYYY') ===
              $root.moment(z).format('DD MMM YYYY')
                ? $style.today
                : '',
              $root.moment(date).format('DD MMM YYYY') === $root.moment(z).format('DD MMM YYYY')
                ? $style.selected
                : '',
              z === null ? $style.null : '',
            ]"
            :style="{ background: getPower(map[$root.moment(z).format('YYYY-MM-DD')]?.calory) }"
            v-for="z in days[i]"
            :key="z"
            :title="map[$root.moment(z).format('YYYY-MM-DD')]?.calory"
          ></div>
        </div>
      </div>
    </div>
    <div style="margin-top: 15px; color: #979797">
      Date: {{ $root.moment(date).format('DD MMM YYYY') }}
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { RestApi } from '../util/RestApi';
import Moment from 'moment';

export default defineComponent({
  props: {
    item: Object,
  },
  components: {},
  async mounted() {
    for (let i = 0; i < 12; i++) {
      const l = this.getDates(i);
      const ll = [];
      for (let j = 0; j < l[0].getDay() - 1; j++) {
        ll.push(null);
      }
      this.days[i] = [...ll, ...l];
    }

    this.getYearMap();
  },
  methods: {
    getPower(calory: number) {
      if (calory === undefined) return 'transparent';
      if (calory <= 0) return '#e2e2e223';
      if (calory >= 2500) return '#e8543c66';
      else if (calory >= 2000) return '#ff930cb0';
      else if (calory >= 1600) return '#fff814a3';
      else return '#26f33f73';
    },
    async getYearMap() {
      this.map = await RestApi.eat.getYearMap(Moment(this.date).format('YYYY-MM-DD'));
    },
    getDates(month: number) {
      const cFrom = new Date();
      const cTo = new Date();

      cFrom.setMonth(month);
      cTo.setMonth(month);
      cFrom.setDate(1);
      cTo.setDate(31);

      let daysArr = [new Date(cFrom)];
      let tempDate = cFrom;

      while (tempDate < cTo) {
        tempDate.setUTCDate(tempDate.getUTCDate() + 1);
        if (tempDate.getMonth() != month) {
          break;
        }
        daysArr.push(new Date(tempDate));
      }

      return daysArr;
    },
  },
  data: () => {
    return {
      month: ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'],
      days: [] as any[],
      date: new Date(),
      map: {},
    };
  },
});
</script>

<style lang="scss" module>
.schedule {
  display: flex;
  flex-direction: column;

  .year {
    display: grid;
    grid-template-columns: 1fr 1fr 1fr 1fr;
    gap: 10px;

    .month {
      flex: 1;
      font-size: 12px;
      color: #acacac;

      .cell {
        width: 12px;
        height: 12px;
        background: #2b2b2b;
        // border: 1px solid transparent;
        font-size: 8px;
        display: flex;
        align-items: center;
        justify-content: center;
        margin-bottom: 2px;

        &.null {
          pointer-events: none;
          background: transparent;
        }

        &.selected {
          border: 1px solid #ca6f00;
        }

        &.today::before {
          content: '●';
          display: flex;
          align-items: center;
          justify-content: center;
          color: #fefefe;
        }
      }
    }
  }
}
</style>
