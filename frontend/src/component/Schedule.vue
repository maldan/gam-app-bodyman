<template>
  <div :class="$style.schedule">
    <div :class="$style.year">
      <div :class="$style.month" v-for="(x, i) in month" :key="x">
        <div>{{ x }}</div>
        <div
          style="
            display: grid;
            margin-top: 10px;
            grid-template-columns: 14px 14px 14px 14px 14px 14px 14px;
          "
        >
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
            ]"
            v-for="z in days[i]"
            :key="z"
          ></div>
          <!-- <div v-for="y in 4" :key="y">
            <div :class="$style.cell" v-for="z in 7" :key="z"></div>
          </div> -->
        </div>
      </div>
    </div>
    <div style="margin-top: 15px">Date: {{ $root.moment(date).format('DD MMM YYYY') }}</div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  props: {
    item: Object,
  },
  components: {},
  async mounted() {
    for (let i = 0; i < 12; i++) {
      this.days[i] = this.getDates(i);
    }
  },
  methods: {
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

      .cell {
        width: 12px;
        height: 12px;
        background: #2b2b2b;
        border: 1px solid #444444;
        font-size: 8px;
        display: flex;
        align-items: center;
        justify-content: center;

        &.selected {
          border: 1px solid #ca6f00;
        }

        &.today::before {
          content: '●';
          display: flex;
          align-items: center;
          justify-content: center;
        }
      }
    }
  }
}
</style>
