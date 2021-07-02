<template>
  <div class="main">
    <Header />

    <div class="body" style="display: flex; height: calc(100% - 50px)">
      <!-- Eat History -->
      <EatHistory :date="currentDate" />

      <!-- Second -->
      <div class="block_row" style="display: flex; flex-direction: column; width: 260px">
        <div class="block">
          <div class="header">Total Eat</div>
          <div class="body" style="display: flex; flex-direction: column">
            <div class="total_item" v-for="(v, k) in stat" :key="k">
              <div>{{ k }}</div>
              <div style="text-align: right">{{ ~~v }}</div>
            </div>
          </div>
        </div>

        <!-- <div class="block">
          <div class="header">Total Vitamin</div>
          <div class="body component_list">
            <div>A</div>
            <div>B</div>
            <div>C</div>
            <div>D</div>
            <div>E</div>
          </div>
        </div>
        <div class="block">
          <div class="header">Total Mineral</div>
          <div class="body component_list">
            <div>Iron</div>
            <div>Iron</div>
            <div>Iron</div>
            <div>Iron</div>
            <div>Iron</div>
          </div>
        </div> -->
      </div>

      <!-- Third -->
      <div style="flex: 1">
        <!-- Schedule -->
        <div class="block">
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
import { RestApi } from '../util/RestApi';
import Eat from '../component/eat/Eat.vue';
import Schedule from '../component/Schedule.vue';
import Header from '../component/Header.vue';
import EatHistory from '../component/eat/History.vue';
import Moment from 'moment';

export default defineComponent({
  components: { Eat, Schedule, Header, EatHistory },
  async mounted() {
    this.refresh();
  },
  methods: {
    async refresh() {
      //this.eat = await RestApi.eat.getFilterByDate(Moment(this.currentDate).format('YYYY-MM-DD'));
      this.stat = await RestApi.eat.getTotalStatByDate(
        Moment(this.currentDate).format('YYYY-MM-DD'),
      );
      (this.$refs['note_list'] as any).refresh();
    },
  },
  data: () => {
    return {
      eatId: '',
      noteId: '',
      //isShowAddEatForm: false,
      //isShowEditEatForm: false,
      isShowAddNoteForm: false,
      isShowEditNoteForm: false,
      //eat: [],
      stat: {},
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

  .body {
    display: grid;
    grid-template-columns: 1fr 1fr 1fr;
    gap: 10px;
    margin-top: 10px;

    .block_row {
      > .block {
        margin-bottom: 10px;
      }
    }
  }

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

  .component_list {
    display: grid;
    grid-gap: 10px;
    grid-template-columns: 1fr 1fr 1fr;

    > div {
      background: #80808045;
      font-size: 14px;
      padding: 5px 10px;
      text-align: center;
      color: #a5a5a5;
      border-radius: 4px;
    }
  }
}
</style>
