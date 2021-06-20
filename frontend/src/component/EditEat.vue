<template>
  <div :class="$style.container">
    <div :class="$style.window">
      <div :class="$style.found_item" v-if="productId">
        {{ productName }}
        <img @click="productId = ''" class="clickable" src="../asset/add.svg" />
      </div>
      <Input
        v-if="!productId"
        icon="add"
        placeholder="Product..."
        style="margin-bottom: 10px"
        v-model="productName"
        @change="searchProduct(form)"
      />
      <Input icon="add" placeholder="Amount..." style="margin-bottom: 10px" v-model="amount" />

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
import Moment from 'moment';

export default defineComponent({
  props: {
    id: String,
  },
  components: { Button, Input, Select },
  async mounted() {
    this.productList = (await RestApi.product.getList()).map((x: any) => {
      return {
        label: x.name,
        value: x.id,
      };
    });

    const eat = await RestApi.eat.get(this.id + '');
    this.productName = eat.product.name;
    this.productId = eat.productId;
    this.amount = eat.amount;
    this.created = Moment(eat.created).format('YYYY-MM-DD HH:mm:ss');
  },
  methods: {
    async searchProduct(form: any) {
      try {
        const product = await RestApi.product.findByName(form.name);
        this.productId = product.id;
        this.productName = product.name;
      } catch {
        try {
          const product = await RestApi.product.findByName(this.defuck(form.name));
          this.productId = product.id;
          this.productName = product.name;
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
      await RestApi.eat.update(this.id + '', this.productId, this.amount, this.created);
      this.$emit('close');
    },
  },
  data() {
    const moment = (this.$root as any)['moment'];

    return {
      productList: [],
      productId: '',
      productName: '',
      amount: '',
      created: moment(new Date()).format('YYYY-MM-DD HH:mm:ss'),
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
