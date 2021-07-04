<template>
  <div :class="$style.container">
    <div :class="$style.window">
      <Input
        v-if="!exerciseId"
        icon="title"
        placeholder="Title..."
        style="margin-bottom: 10px"
        v-model="title"
        @change="searchExercise(title)"
      />

      <div :class="$style.found_item" v-if="exerciseId">
        {{ title }}
        <img @click="exerciseId = ''" class="clickable" src="../../asset/remove.svg" />
      </div>

      <!-- <Select
        icon="date"
        placeholder="Tool..."
        style="margin-bottom: 10px"
        :items="toolList"
        v-model="tool"
      /> -->

      <Input icon="date" placeholder="Reps..." style="margin-bottom: 10px" v-model="reps" />
      <Input icon="weight" placeholder="Weight..." style="margin-bottom: 10px" v-model="weight" />
      <Input
        icon="distance"
        placeholder="Distance..."
        style="margin-bottom: 10px"
        v-model="distance"
      />
      <Input icon="time" placeholder="Duration..." style="margin-bottom: 10px" v-model="duration" />
      <Input
        icon="date"
        placeholder="Date..."
        style="margin-bottom: 10px"
        v-model="created"
        functionIcon="date"
        :functionClick="currentTime"
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
import { Helper } from '../../util/Helper';

export default defineComponent({
  props: {
    id: String,
    date: Object,
  },
  components: { Button, Input, Select },
  async mounted() {
    if (this.id) {
      const training = await RestApi.training.get(this.id + '');

      this.exerciseId = training.exerciseId;
      this.title = training.title;
      // this.title = training.title;
      //  this.tool = { label: training.tool, value: training.tool };
      this.reps = training.reps;
      this.distance = training.distance;
      this.weight = training.weight;
      this.duration = Moment.utc(
        Moment.duration(training.duration, 'seconds').asMilliseconds(),
      ).format('HH:mm:ss');
      this.created = Moment(training.created).format('YYYY-MM-DD HH:mm:ss');
    }
  },
  methods: {
    async searchExercise() {
      try {
        const product = await RestApi.exercise.findByName(this.title);
        this.exerciseId = product.id;
        this.title = product.name;
      } catch {
        try {
          const product = await RestApi.exercise.findByName(Helper.traslit(this.title));
          this.exerciseId = product.id;
          this.title = product.name;
        } catch {}
      }
    },
    async submit() {
      await RestApi.training.add(
        this.exerciseId,
        Number(this.reps),
        Number(this.weight),
        Number(this.distance),
        Moment.duration(this.duration).asSeconds(),
        this.created,
      );
      this.$emit('close');
    },
    currentTime() {
      return Moment().format('YYYY-MM-DD HH:mm:ss');
    },
  },
  data() {
    const moment = (this.$root as any)['moment'];
    const d = this.date as Date;
    d.setHours(new Date().getHours());
    d.setMinutes(new Date().getMinutes());
    d.setSeconds(new Date().getSeconds());

    return {
      /*toolList: [
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
      title: '',*/

      title: '',
      exerciseId: '',
      distance: '',
      reps: '',
      weight: '',
      duration: '00:00:00',
      created: moment(d).format('YYYY-MM-DD HH:mm:ss'),
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
