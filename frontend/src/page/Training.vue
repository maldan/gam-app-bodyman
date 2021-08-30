<template>
  <div class="main">
    <Header />

    <div style="display: flex; height: calc(100% - 60px); margin-top: 10px">
      <History :date="currentDate" />

      <!-- Second -->
      <div style="display: flex; flex-direction: column; width: 290px; margin-right: 10px">
        <div class="block" style="margin-bottom: 10px">
          <div class="header">Total Power</div>
          <div class="body">
            <div class="total_item" v-for="(v, k) in stat.tool" :key="k">
              <div @click="(type = 'tool'), (mode = k)" class="clickable">
                {{ k.split('_').join(' ') }}
              </div>
              <div style="text-align: right">{{ ~~v }}</div>
            </div>
          </div>
        </div>

        <div class="block" style="margin-bottom: 10px">
          <div class="header">Total Power</div>
          <div class="body">
            <div class="total_item" v-for="(v, k) in stat.muscle" :key="k">
              <div @click="(type = 'muscle'), (mode = k)" class="clickable">
                {{ k.split('_').join(' ') }}
              </div>
              <div style="text-align: right">{{ ~~v }}</div>
            </div>
          </div>
        </div>
      </div>

      <!-- Third -->
      <div style="flex: 1">
        <!-- Schedule -->
        <div class="block" style="margin-bottom: 10px">
          <div class="header">Schedule</div>
          <div class="body">
            <Schedule :type="type" :mode="mode" @select="(currentDate = $event), refresh()" />
          </div>
        </div>

        <!-- Schedule -->
        <Fap :date="currentDate" style="margin-top: 10px" />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import History from '../component/training/History.vue';
import Fap from '../component/training/Fap.vue';
import Header from '../component/Header.vue';
import Schedule from '../component/training/Schedule.vue';
import { RestApi } from '../util/RestApi';
import Moment from 'moment';

export default defineComponent({
  components: { Header, History, Schedule, Fap },
  async mounted() {
    this.refresh();
  },
  methods: {
    async refresh() {
      this.stat = await RestApi.training.getTotalStatByDate(
        Moment(this.currentDate).format('YYYY-MM-DD'),
      );
    },
  },
  data: () => {
    return {
      currentDate: new Date(),
      stat: {},
      type: 'tool',
      mode: 'total',
    };
  },
});
</script>

<style lang="scss" scoped>
.main {
  box-sizing: border-box;
  height: 100%;
  padding: 10px;

  .body {
    .total_item {
      display: flex;
      background: #80808045;
      font-size: 15px;
      color: #a5a5a5;
      text-transform: capitalize;
      margin-bottom: 5px;
      padding: 5px 10px;
      border-radius: 4px;

      > div {
        flex: 1;
      }

      > div:last-child {
        font-weight: bold;
      }
    }
  }
}
</style>
