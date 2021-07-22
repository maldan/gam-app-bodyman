<template>
  <div class="main">
    <Header />

    <Button
      @click="isShowAddForm = !isShowAddForm"
      text="Add"
      icon="add"
      style="flex: 0; margin-top: 10px"
    />

    <List
      style="margin-top: 10px"
      ref="note_list"
      @edit="(isShowEditForm = true), (noteId = $event)"
    />

    <Add :date="currentDate" v-if="isShowAddForm" @close="(isShowAddForm = false), refresh()" />
    <Edit :id="noteId" v-if="isShowEditForm" @close="(isShowEditForm = false), refresh()" />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { RestApi } from '../util/RestApi';
import Add from '../component/note/Add.vue';
import Edit from '../component/note/Edit.vue';
import List from '../component/note/List.vue';
import Header from '../component/Header.vue';
import Button from '../component/Button.vue';
import Moment from 'moment';

export default defineComponent({
  components: { List, Add, Edit, Header, Button },
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
      isShowAddForm: false,
      isShowEditForm: false,
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
  display: flex;
  flex-direction: column;
}
</style>
