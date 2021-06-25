<template>
  <div class="block" style="margin-right: 10px; width: 320px; height: calc(100% - 10px)">
    <!-- Header -->
    <div class="header">
      Eat history
      <img @click="isShowAddForm = true" class="clickable" src="../../asset/add.svg" alt="" />
    </div>

    <!-- List -->
    <div class="body" style="height: calc(100% - 22px - 15px)">
      <Eat
        @edit="(isShowEditForm = true), (eatId = $event)"
        v-for="(x, i) in eatList"
        :key="x.id"
        :item="x"
        :nextItem="eatList[i + 1]"
        :date="date"
        style="margin-bottom: 15px"
      />
    </div>

    <Add :date="date" v-if="isShowAddForm" @close="(isShowAddForm = false), refresh()" />
    <Edit :id="eatId" v-if="isShowEditForm" @close="(isShowEditForm = false), refresh()" />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { RestApi } from '../../util/RestApi';
import Eat from './Eat.vue';
import Add from './Add.vue';
import Edit from './Edit.vue';
import Moment from 'moment';

export default defineComponent({
  props: {
    date: Object,
  },
  components: { Eat, Add, Edit },
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
      this.eatList = await RestApi.eat.getFilterByDate(Moment(this.date).format('YYYY-MM-DD'));
    },
  },
  data: () => {
    return {
      eatId: '',
      isShowAddForm: false,
      isShowEditForm: false,
      eatList: [],
    };
  },
});
</script>

<style lang="scss" scoped></style>
