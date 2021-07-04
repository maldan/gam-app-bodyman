<template>
  <div :class="$style.note_list">
    <div v-for="note in noteList" :key="note.id" :class="$style.block">
      <div :class="$style.header">
        <img
          @click="$emit('edit', note.id)"
          class="clickable"
          src="../../asset/pencil.svg"
          alt=""
          style="margin-left: auto"
        />
        <img @click="remove(note.id)" class="clickable" src="../../asset/trash.svg" alt="" />
        <div :class="$style.right">{{ $root.moment(note.created).fromNow() }}</div>
      </div>
      <div :class="$style.body">
        {{ note.description }}
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { RestApi } from '../../util/RestApi';

export default defineComponent({
  props: {},
  components: {},
  async mounted() {
    this.refresh();
  },
  methods: {
    async refresh() {
      this.noteList = await RestApi.note.getList();
    },
    async remove(id: string) {
      if (confirm('Are you sure?')) {
        await RestApi.note.delete(id);
      }
      this.refresh();
    },
  },
  data: () => {
    return {
      noteList: [],
    };
  },
});
</script>

<style lang="scss" module>
.note_list {
  .block {
    font-size: 15px;
    margin-bottom: 15px;

    .header {
      display: flex;

      .left,
      .right {
        padding: 5px 10px;
        background: #0000004d;
        border-radius: 6px 6px 0 0;
        color: #979797;
        font-weight: bold;

        span {
          color: #979797;
          font-style: italic;
          font-weight: 300;
        }
      }

      .left {
        color: #7be01e;
      }

      img {
        margin-left: 15px;
      }

      .right {
        margin-left: 15px;
      }
    }

    .body {
      padding: 10px 15px;
      background: #0000004d;
      border-radius: 0 0 6px 6px;
      color: #979797;
    }

    &:last-child {
      margin-bottom: 0;
    }
  }
}
</style>
