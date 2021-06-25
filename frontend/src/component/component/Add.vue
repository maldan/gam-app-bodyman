<template>
  <div :class="$style.container">
    <div :class="$style.window">
      <Input icon="food" placeholder="Name..." style="margin-bottom: 10px" v-model="name" />
      <Select
        icon="food"
        placeholder="Type..."
        style="margin-bottom: 10px"
        v-model="type"
        :items="[
          { label: 'Vitamin', value: 'vitamin' },
          { label: 'Mineral', value: 'mineral' },
          { label: 'Other', value: 'other' },
        ]"
      />

      <div style="display: flex">
        <Button @click="$emit('close')" text="Cancel" style="margin-right: 5px" />
        <Button @click="submit()" text="Add" icon="add" style="margin-left: 5px" />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { RestApi } from '../../util/RestApi';
import Button from '../Button.vue';
import Input from '../Input.vue';
import Select from '../Select.vue';

export default defineComponent({
  props: {
    date: Object,
  },
  components: { Button, Input, Select },
  async mounted() {},
  methods: {
    async submit() {
      await RestApi.component.add(this.name, this.type.value);
      this.$emit('close');
    },
  },
  data() {
    return {
      name: '',
      type: {
        label: '',
        value: '',
      },
    };
  },
});
</script>

<style lang="scss" module>
.container {
  position: fixed;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.6);
  z-index: 10;
  display: flex;
  align-items: center;
  justify-content: center;

  .window {
    min-width: 320px;
    background: #444444;
    border-radius: 4px;
    padding: 25px;
    color: #fefefe;
    box-shadow: 0px 1px 6px 2px #00000055;

    .found_item {
      padding: 10px 15px;
      background: #3a3a3a;
      margin-bottom: 10px;
      border-radius: 4px;
      display: flex;

      img {
        margin-left: auto;
      }
    }
  }
}
</style>
