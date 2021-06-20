<template>
  <div :class="$style.container">
    <div :class="$style.window">
      <div v-for="form in foodForm" :key="form">
        <div :class="$style.found_item" v-if="form.id">
          {{ form.name }}
          <img @click="form.id = ''" class="clickable" src="../asset/add.svg" />
        </div>
        <Input
          v-if="!form.id"
          icon="add"
          placeholder="Product..."
          style="margin-bottom: 10px"
          v-model="form.name"
          @change="searchProduct(form)"
        />
        <Input
          icon="add"
          placeholder="Amount..."
          style="margin-bottom: 10px"
          v-model="form.amount"
        />
      </div>

      <Input icon="add" placeholder="Date..." style="margin-bottom: 10px" v-model="created" />

      <div style="display: flex">
        <Button @click="$emit('close')" text="Cancel" style="margin-right: 5px" />
        <Button @click="submit()" text="Add" icon="add" style="margin-left: 5px" />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { RestApi } from '../util/RestApi';
import Button from './Button.vue';
import Input from './Input.vue';
import Select from './Select.vue';

export default defineComponent({
  props: {
    date: Object,
  },
  components: { Button, Input, Select },
  async mounted() {
    for (let i = 0; i < 4; i++) {
      this.foodForm.push({
        id: '',
        name: '',
        amount: '',
      });
    }
    this.productList = (await RestApi.product.getList()).map((x: any) => {
      return {
        label: x.name,
        value: x.id,
      };
    });
  },
  methods: {
    async searchProduct(form: any) {
      try {
        const product = await RestApi.product.findByName(form.name);
        form.id = product.id;
        form.name = product.name;
      } catch {
        try {
          const product = await RestApi.product.findByName(this.defuck(form.name));
          form.id = product.id;
          form.name = product.name;
        } catch {}
      }
    },
    defuck(x: string) {
      const a = "`qwertyuiop[]asdfghjkl;'zxcvbnm,.";
      const b = 'ёйцукенгшщзхъфывапролджэячсмитьбю';
      let out = x.split('');
      for (let i = 0; i < x.length; i++) {
        for (let j = 0; j < a.length; j++) {
          if (x[i] === a[j]) {
            out[i] = b[j];
            break;
          }
        }
      }
      return out.join('');
    },
    async submit() {
      for (let i = 0; i < this.foodForm.length; i++) {
        if (!this.foodForm[i].id) {
          continue;
        }
        await RestApi.eat.add(this.foodForm[i].id, this.foodForm[i].amount, this.created);
      }
      this.$emit('close');
    },
  },
  data() {
    const moment = (this.$root as any)['moment'];

    return {
      productList: [],
      created: moment(this.date).format('YYYY-MM-DD HH:mm:ss'),
      foodForm: [] as any[],
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
