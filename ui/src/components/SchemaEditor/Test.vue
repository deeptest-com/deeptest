<script setup>
import {ref} from 'vue';
import {computePosition} from '@floating-ui/dom';


const reference = ref(null);
const floating = ref(null);

const x = ref(0);
const y = ref(0);
const strategy = ref('');


// const {x, y, strategy} = useFloating(reference, floating);

function click(e) {
  computePosition(e.target, floating.value).then(({x, y}) => {
    Object.assign(floating.value.style, {
      left: `${x}px`,
      top: `${y}px`,
    });
  });
}

</script>

<template>
  <button @click="click">Button</button>
  <button @click="click">Button1</button>
  <button @click="click">Button2</button>
  <button @click="click">Button3</button>
  <div
      :style="{
      position: 'fixed',
      top: `${y ?? 0}px`,
      left: `${x ?? 0}px`,
    }"
      ref="floating"
  >
    Tooltip
  </div>
</template>
