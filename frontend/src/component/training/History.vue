<template>
  <div class="block" style="margin-right: 10px; width: 320px; height: calc(100% - 10px)">
    <!-- Header -->
    <div class="header">
      Training history
      <img @click="isShowAddForm = true" class="clickable" src="../../asset/add.svg" alt="" />
    </div>

    <!-- List -->
    <div class="body" style="height: calc(100% - 22px - 15px)">
      <Training
        @edit="(isShowEditForm = true), (id = $event)"
        @copy="(isShowAddForm = true), (id = $event)"
        @delete="remove(x.id)"
        v-for="(x, i) in list"
        :key="x.id"
        :item="x"
        :nextItem="list[i + 1]"
        :date="date"
        style="margin-bottom: 15px"
      />
    </div>

    <Add
      :id="id"
      :date="date"
      v-if="isShowAddForm"
      @close="((isShowAddForm = false), (id = '')), refresh()"
    />
    <Edit
      :id="id"
      v-if="isShowEditForm"
      @close="((isShowEditForm = false), (id = '')), refresh()"
    />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { RestApi } from '../../util/RestApi';
import Training from './Training.vue';
import Add from './Add.vue';
import Edit from './Edit.vue';
import Moment from 'moment';

export default defineComponent({
  props: {
    date: Object,
  },
  components: { Training, Add, Edit },
  async mounted() {
    this.refresh();
  },
  watch: {
    date(value: Date) {
      this.refresh();
    },
  },
  methods: {
    async refresh() {
      this.list = await RestApi.training.getFilterByDate(Moment(this.date).format('YYYY-MM-DD'));
    },
    async remove(id: string) {
      if (confirm('Are you sure?')) {
        await RestApi.training.delete(id);
      }
      this.refresh();
    },
  },
  data: () => {
    return {
      id: '',
      isShowAddForm: false,
      isShowEditForm: false,
      list: [],
    };
  },
});
</script>

<style lang="scss" scoped></style>
