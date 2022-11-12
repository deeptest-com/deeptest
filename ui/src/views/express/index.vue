<template>
  <div id="express-main" class="express-main dp-splits-v">
        <div>
          <md-editor v-model="content" class="dp-md" preview-only />
        </div>

<!--    <div id="express-left" class="left">-->
<!--      <InterfaceTree />-->
<!--    </div>-->

<!--    <div id="express-splitter" class="splitter"></div>-->

<!--    <div id="express-right" class="right">-->
<!--      <InterfaceDesigner></InterfaceDesigner>-->
<!--    </div>-->
  </div>
</template>

<script setup lang="ts">
import {computed, defineComponent, onMounted, onUnmounted, Ref, ref, watch} from "vue";
import MdEditor from 'md-editor-v3';
import 'md-editor-v3/lib/style.css';

// import InterfaceDesigner from './components/Designer.vue';
import {resizeWidth} from "@/utils/dom";
// import InterfaceTree from "./components/Tree.vue"
import {useStore} from "vuex";
import {StateType as GlobalStateType} from "@/store/global";
import {StateType as UserStateType} from "@/store/user";
import {StateType as ProjectStateType} from "@/store/project";
import {Interface} from "@/views/interface/data";
import {StateType} from "@/views/interface/store";
import {StateType as SpecStateType} from "@/views/express/store";

const store = useStore<{ Global: GlobalStateType, Spec: SpecStateType}>();
const specData = computed<any>(() => store.state.Spec.specData);
const content = ref('')

onMounted(() => {
  console.log('onMounted')
  resize()
})
onUnmounted(() => {
  console.log('onUnmounted')
})

watch(specData, () => {
  console.log('watch specData')
  content.value = specData.value.doc?.info?.description
}, {deep: true})

const resize = () => {
  resizeWidth('express-main',
        'express-left', 'express-splitter', 'express-right',
      260, 800)
}

</script>

<style lang="less">
.dp-md {
  h1 {
    margin: 6px 0px 3px 0px !important;
  }
}
</style>

<style lang="less" scoped>
.express-main {
  .left {
    width: 260px;
  }
  .right {
    flex: 1;
  }

  .dp-md {
    h1 {
      margin: 6px 0px 6px 0px !important;
    }
  }
}
</style>