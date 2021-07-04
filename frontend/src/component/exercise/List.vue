<template>
  <div class="block" style="flex: 1; height: calc(100% - 10px)">
    <!-- Header -->
    <div class="header">
      Exercise List
      <img @click="isShowAddForm = true" class="clickable" src="../../asset/add.svg" alt="" />
    </div>

    <!-- List -->
    <div class="body" style="height: calc(100% - 22px - 15px)">
      <Input placeholder="Filter..." style="margin-bottom: 10px" v-model="filter" />

      <div
        class="product_item"
        v-for="x in list
          .filter(
            (x) =>
              x.name.toLowerCase().match(filter.toLowerCase()) ||
              x.name.toLowerCase().match(translit(filter.toLowerCase())),
          )
          .sort((a, b) => a.name.localeCompare(b.name))"
        :key="x.id"
      >
        <div class="product_header">
          <div>{{ x.name }}</div>
          <img
            @click="(isShowEditForm = true), (productId = x.id)"
            class="clickable"
            src="../../asset/pencil.svg"
            alt=""
            style="margin-left: auto"
          />
          <img @click="remove(x.id)" class="clickable" src="../../asset/trash.svg" alt="" />
        </div>
        <div class="product_component">
          <div v-for="y in x.muscleList" :key="y">
            <span>{{ y }}</span>
          </div>
        </div>
      </div>

      <Add v-if="isShowAddForm" @close="(isShowAddForm = false), refresh()" />
      <Edit :id="productId" v-if="isShowEditForm" @close="(isShowEditForm = false), refresh()" />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { RestApi } from '../../util/RestApi';
import Header from '../Header.vue';
import Input from '../Input.vue';
import Add from './Add.vue';
import Edit from './Edit.vue';
import { Helper } from '../../util/Helper';

export default defineComponent({
  components: { Header, Input, Add, Edit },
  async mounted() {
    this.refresh();
  },
  methods: {
    translit: Helper.traslit,
    async refresh() {
      this.list = (await RestApi.exercise.getList()) || [];
    },
    async remove(id: string) {
      if (confirm('Are you sure?')) {
        if (confirm('Are you really really sure?')) {
          await RestApi.exercise.delete(id);
        }
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
      filter: '',
    };
  },
});
</script>

<style lang="scss" scoped>
.main {
  box-sizing: border-box;
  height: 100%;
  padding: 10px;

  .product_item {
    padding: 10px;
    background: #383838;
    color: #b1b1b1;
    border-radius: 4px;
    margin-bottom: 10px;

    .product_header {
      display: flex;
      align-items: flex-start;

      img {
        margin-left: 10px;
      }
    }

    .product_component {
      display: grid;
      grid-template-columns: 1fr 1fr 1fr;
      gap: 10px;
      font-size: 14px;
      margin-top: 10px;

      > div {
        display: flex;
        background: #292929;
        padding: 2px 5px;
        text-align: center;
        color: #929292;
        border-radius: 4px;

        span {
          color: #6e6e6e;
          margin-right: auto;
        }
      }
    }
  }
}
</style>
