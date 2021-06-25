<template>
  <div :class="$style.block">
    <div :class="$style.header">
      <div :class="$style.left">
        {{ item.reps ? `${item.weight}x${item.reps}` : `${item.distance} m` }}
        <span>{{ item.tool }}</span>
      </div>
      <div :class="$style.right">{{ $root.moment(item.created).format('HH:mm') }}</div>
    </div>
    <div :class="$style.body">
      <div :class="$style.name">
        {{ item.title }}
        <img
          @click="$emit('copy', item.id)"
          class="clickable"
          src="../../asset/copy.svg"
          alt=""
          style="margin-left: auto"
        />
        <img
          @click="$emit('edit', item.id)"
          class="clickable"
          src="../../asset/pencil.svg"
          alt=""
        />
        <img
          @click="$emit('delete', item.id)"
          class="clickable"
          src="../../asset/trash.svg"
          alt=""
        />
      </div>
      <div style="margin-top: 5px">
        {{
          $root.moment
            .utc($root.moment.duration(item.duration, 'seconds').asMilliseconds())
            .format('HH:mm:ss')
        }}
      </div>
      <!-- <div :class="$style.component">
        <div><span>P</span> {{ item.product.protein }}</div>
        <div><span>C</span> {{ item.product.carbohydrate }}</div>
        <div><span>F</span> {{ item.product.fat }}</div>
      </div> -->
    </div>
    <div :class="$style.gap" v-if="gap()">{{ gap() }}</div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import Moment from 'moment';

export default defineComponent({
  props: {
    item: Object,
    nextItem: Object,
    date: Object,
  },
  components: {},
  async mounted() {},
  methods: {
    gap() {
      if (!this.item) {
        return null;
      }
      if (
        !this.nextItem &&
        Moment(this.date).format('YYYY-MM-DD') !== Moment(new Date()).format('YYYY-MM-DD')
      ) {
        return null;
      }

      const a = new Date(this.item.created).getTime() / 1000;
      const b = new Date(this.nextItem?.created || new Date()).getTime() / 1000;
      return Moment.utc(Moment.duration(b - a, 'seconds').asMilliseconds()).format('HH:mm:ss');
    },
  },
  data: () => {
    return {};
  },
});
</script>

<style lang="scss" module>
.block {
  font-size: 15px;

  .header {
    display: flex;

    .left,
    .right {
      padding: 5px 10px;
      border-radius: 6px 6px 0 0;
      color: #b1b1b1;
      background: #2c2c2c;
      font-weight: bold;

      span {
        color: #bdbdbd;
        font-style: italic;
        font-weight: 300;
      }
    }

    .left {
      color: #7be01e;
    }

    .right {
      margin-left: auto;
    }
  }

  .body {
    padding: 10px 15px;
    background: #383838;
    border-radius: 0 0 6px 6px;
    color: #b1b1b1;

    .name {
      display: flex;
      color: #ffe10c;

      img {
        margin-left: 15px;
      }
    }

    .component {
      margin-top: 10px;
      display: grid;
      align-items: center;
      grid-template-columns: 1fr 1fr 1fr;
      gap: 5px;
      font-size: 14px;

      > div {
        display: flex;

        span {
          color: #828282;
          margin-right: auto;
        }
        flex: 1;
        background: #292929;
        padding: 5px 10px;
      }
    }
  }

  .gap {
    background: #2c2c2c;
    padding: 3px 10px;
    text-align: center;
    border-radius: 4px;
    margin-top: 20px;
    color: #7be01e;
    position: relative;
    width: max-content;
    margin-left: auto;
    margin-right: auto;

    &::before {
      content: '|';
      position: absolute;
      left: 50%;
      top: -22px;
      color: #828282;
    }

    &::after {
      content: '|';
      position: absolute;
      left: 50%;
      bottom: -20px;
      color: #828282;
    }
  }
}
</style>
