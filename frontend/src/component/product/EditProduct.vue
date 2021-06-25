<template>
  <div :class="$style.container">
    <div :class="$style.window">
      <Input icon="food" placeholder="Product..." style="margin-bottom: 10px" v-model="name" />
      <Input icon="weight" placeholder="Protein..." style="margin-bottom: 10px" v-model="protein" />
      <Input
        icon="weight"
        placeholder="Carbohydrate..."
        style="margin-bottom: 10px"
        v-model="carbohydrate"
      />
      <Input icon="weight" placeholder="Fat..." style="margin-bottom: 10px" v-model="fat" />

      <div style="display: flex">
        <Button @click="$emit('close')" text="Cancel" style="margin-right: 5px" />
        <Button @click="submit()" text="Save" icon="pencil" style="margin-left: 5px" />
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
    id: String,
  },
  components: { Button, Input, Select },
  async mounted() {
    const product = await RestApi.product.get(this.id + '');
    this.name = product.name;
    this.protein = product.protein;
    this.carbohydrate = product.carbohydrate;
    this.fat = product.fat;
  },
  methods: {
    async submit() {
      await RestApi.product.update(
        this.id + '',
        this.name,
        this.protein,
        this.carbohydrate,
        this.fat,
      );
      this.$emit('close');
    },
  },
  data() {
    return {
      name: '',
      protein: '',
      carbohydrate: '',
      fat: '',
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
