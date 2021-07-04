<template>
  <div :class="$style.container">
    <div :class="$style.window">
      <TextArea placeholder="Description..." style="margin-bottom: 10px" v-model="description" />

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
import TextArea from '../TextArea.vue';

export default defineComponent({
  props: {
    date: Object,
  },
  components: { Button, TextArea },
  async mounted() {},
  methods: {
    async submit() {
      await RestApi.note.add(this.description);
      this.$emit('close');
    },
  },
  data() {
    return {
      description: '',
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
  }
}
</style>
