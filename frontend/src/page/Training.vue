<template>
  <div class="main">
    <Header />

    <div style="display: flex; height: calc(100% - 60px); margin-top: 10px">
      <History :date="currentDate" />

      <!-- Second -->
      <div style="display: flex; flex-direction: column; width: 260px; margin-right: 10px">
        <div class="block" style="margin-bottom: 10px">
          <div class="header">Total Power</div>
          <div class="body">
            <div class="total_item" v-for="(v, k) in stat" :key="k">
              <div>{{ k }}</div>
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
            <Schedule @select="(currentDate = $event), refresh()" />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import History from '../component/training/History.vue';
import Header from '../component/Header.vue';
import Schedule from '../component/Schedule.vue';

export default defineComponent({
  components: { Header, History, Schedule },
  mounted() {},
  methods: {
    refresh() {},
  },
  data: () => {
    return {
      currentDate: new Date(),
    };
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
