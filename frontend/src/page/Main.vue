<template>
  <div class="main">
    <div style="display: flex; height: 100%">
      <!-- Eat History -->
      <div class="block" style="margin-right: 15px; width: 320px; height: 100%">
        <!-- Header -->
        <div class="header">
          Eat history
          <img @click="isShowAddEatForm = true" class="clickable" src="../asset/add.svg" alt="" />
        </div>

        <!-- List -->
        <div class="body" style="height: calc(100% - 22px - 15px)">
          <Eat
            @edit="(isShowEditEatForm = true), (eatId = $event)"
            v-for="x in eat"
            :key="x.id"
            :item="x"
            style="margin-bottom: 15px"
          />
        </div>
      </div>

      <!-- Second -->
      <div style="display: flex; flex-direction: column; width: 260px; margin-right: 15px">
        <div class="block" style="margin-bottom: 15px">
          <div class="header">Total Eat</div>
          <div class="body">
            <div class="total_item" v-for="(v, k) in stat" :key="k">
              <div>{{ k }}</div>
              <div style="text-align: right">{{ ~~v }}</div>
            </div>
          </div>
        </div>

        <div class="block" style="margin-bottom: 15px">
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
        </div>
      </div>

      <!-- Third -->
      <div style="flex: 1">
        <!-- Schedule -->
        <div class="block" style="margin-bottom: 15px">
          <div class="header">Schedule</div>
          <div class="body">
            <Schedule @select="(currentDate = $event), refresh()" />
          </div>
        </div>

        <!-- Notes -->
        <div class="block" style="height: calc(100% - 420px)">
          <div class="header">
            Notes
            <img
              @click="isShowAddNoteForm = true"
              class="clickable"
              src="../asset/add.svg"
              alt=""
            />
          </div>
          <div class="body" style="height: calc(100% - 22px - 15px)">
            <NoteList ref="note_list" @edit="(isShowEditNoteForm = true), (noteId = $event)" />
          </div>
        </div>
      </div>
    </div>

    <AddNote
      :date="currentDate"
      v-if="isShowAddNoteForm"
      @close="(isShowAddNoteForm = false), refresh()"
    />
    <AddEat
      :date="currentDate"
      v-if="isShowAddEatForm"
      @close="(isShowAddEatForm = false), refresh()"
    />
    <EditEat :id="eatId" v-if="isShowEditEatForm" @close="(isShowEditEatForm = false), refresh()" />
    <EditNote
      :id="noteId"
      v-if="isShowEditNoteForm"
      @close="(isShowEditNoteForm = false), refresh()"
    />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { RestApi } from '../util/RestApi';
import Eat from '../component/Eat.vue';
import AddNote from '../component/AddNote.vue';
import AddEat from '../component/AddEat.vue';
import EditEat from '../component/EditEat.vue';
import EditNote from '../component/EditNote.vue';
import Schedule from '../component/Schedule.vue';
import NoteList from '../component/NoteList.vue';
import Moment from 'moment';

export default defineComponent({
  components: { Eat, AddEat, Schedule, EditEat, NoteList, AddNote, EditNote },
  async mounted() {
    this.refresh();
  },
  methods: {
    async refresh() {
      this.eat = await RestApi.eat.getFilterByDate(Moment(this.currentDate).format('YYYY-MM-DD'));
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
      isShowAddEatForm: false,
      isShowEditEatForm: false,
      isShowAddNoteForm: false,
      isShowEditNoteForm: false,
      eat: [],
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

  .total_item {
    display: flex;
    background: #2b2b2b;
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
      background: #2b2b2b;
      font-size: 14px;
      padding: 5px 10px;
      text-align: center;
      color: #a5a5a5;
      border-radius: 4px;
    }
  }
}
</style>
