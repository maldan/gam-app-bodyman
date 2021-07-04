<template>
  <div :class="$style.container">
    <div :class="$style.window">
      <Input icon="title" placeholder="Name..." style="margin-bottom: 10px" v-model="name" />

      <Select
        icon="date"
        placeholder="Tool..."
        style="margin-bottom: 10px"
        :items="toolList"
        v-model="tool"
      />

      <Input
        icon="date"
        placeholder="Muscle List..."
        style="margin-bottom: 10px"
        v-model="muscleList"
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
import Moment from 'moment';

export default defineComponent({
  props: {
    id: String,
  },
  components: { Button, Input, Select },
  async mounted() {},
  methods: {
    async submit() {
      await RestApi.exercise.add(this.name, this.tool.value, this.muscleList.split(', '));
      this.$emit('close');
    },
    currentTime() {
      return Moment().format('YYYY-MM-DD HH:mm:ss');
    },
  },
  data() {
    return {
      toolList: [
        { label: 'Own Weight', value: 'own_weight' },
        { label: 'Barbell', value: 'barbell' },
        { label: 'Dumbbell', value: 'dumbbell' },
        { label: 'Dumbbell x 2', value: 'dumbbell_2' },
        { label: 'Spring Gripper', value: 'spring_gripper' },
        { label: 'Skipping Rope', value: 'skipping_rope' },
      ],
      tool: {
        label: '',
        value: '',
      },
      name: '',
      muscleList: '',
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
