<template>
  <div class="main">
    <Header />

    <NoteList
      style="margin-top: 10px"
      ref="note_list"
      @edit="(isShowEditNoteForm = true), (noteId = $event)"
    />

    <AddNote
      :date="currentDate"
      v-if="isShowAddNoteForm"
      @close="(isShowAddNoteForm = false), refresh()"
    />

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
import AddNote from '../component/AddNote.vue';
import EditNote from '../component/EditNote.vue';
import NoteList from '../component/NoteList.vue';
import Header from '../component/Header.vue';
import Moment from 'moment';

export default defineComponent({
  components: { NoteList, AddNote, EditNote, Header },
  async mounted() {
    this.refresh();
  },
  methods: {
    async refresh() {
      this.stat = await RestApi.eat.getTotalStatByDate(
        Moment(this.currentDate).format('YYYY-MM-DD'),
      );
      (this.$refs['note_list'] as any).refresh();
    },
  },
  data: () => {
    return {
      noteId: '',
      isShowAddNoteForm: false,
      isShowEditNoteForm: false,
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
}
</style>
