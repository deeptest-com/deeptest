<template>
  <!-- 选择环境 -->
  <Teleport to="body" :disabled="disabled">
    <div
      v-if="show"
      :class="`select-env-container ${!disabled ? 'fixed' : ''}`"
      :style="{
        top: selectEnvTopPosition,
        right: selectEnvLeftPosition,
      }"
    >
      <a-select
        :value="currServerId"
        :options="servers"
        @focus="handleFocus"
        @change="e => $emit('change', e)"
        placeholder="请选择环境"
      >
        <template #dropdownRender="{ menuNode: menu }">
          <v-nodes :vnodes="menu" />
          <a-divider style="margin: 4px 0" />
          <a-button type="link" @click="handleRedirectEnv">
            <SettingOutlined />
            环境管理</a-button>
        </template>
      </a-select>
    </div>
  </Teleport>
</template>
<script setup lang="ts">
import {
  defineProps,
  computed,
  defineEmits,
  ref,
  onMounted,
  watch,
  inject,
  defineComponent,
} from "vue";
import { useStore } from "vuex";
import { StateType as Debug } from "@/views/component/debug/store";
import { SettingOutlined } from "@ant-design/icons-vue";
import { useRouter } from "vue-router";

const VNodes = defineComponent({
  props: {
    vnodes: {
      type: Object,
      required: true,
    },
  },
  render() {
    return this.vnodes;
  },
});

const props = defineProps<{
  serverId: any;
  show: boolean;
  disabled: boolean;
}>();

const emits = defineEmits(["change"]);

const containerScrollTop = inject("containerScrollTop", null) as any;
const router = useRouter();

const store = useStore<{ Debug: Debug; Endpoint; Global; ServeGlobal; }>();
const servers = computed<any[]>(() => store.state.Debug.serves);
const currServerId = computed<any[]>(() => store.state.Debug.currServe.environmentId || null);  //当前选择的环境id
const currServe = computed<any>(() => store.state.ServeGlobal.currServe); // 当前选择的服务

const selectEnvTopPosition = ref("0px");
const selectEnvLeftPosition = ref("0px");

onMounted(() => {
  selectEnvTopPosition.value = getSelectEnvTopPosition();
  selectEnvLeftPosition.value = getSelectEnvLeftPosition();
});

const getSelectEnvTopPosition = () => {
  const elems = document.getElementsByClassName("tab-header-items");
  const selectEnvEl = document.getElementsByClassName("select-env-container");
  if (elems.length === 0 || selectEnvEl.length === 0) return "0px";

  const rect = elems[0].getBoundingClientRect();
  const selecEnvRect = selectEnvEl[0].getBoundingClientRect();
  if (!rect || !selecEnvRect) return "0px";

  const { height = 0, top = 0 } = rect;
  const { height: envElHeight = 0 } = selecEnvRect;

  return `${top + (height - envElHeight) / 2}px`;
};

const getSelectEnvLeftPosition = () => {
  const elems = document.getElementsByClassName("tab-header-btns");
  if (elems.length === 0) {
    return "22px";
  }
  if (elems[0].children.length === 0) {
    return "22px";
  }
  const rect = elems[0].getBoundingClientRect();
  if (!rect) {
    return "22px";
  }
  const { width = 0 } = rect;
  return `${width + 16 + 20}px`;
};

const handleRedirectEnv = (e) => {
  e.preventDefault();
  window.open(`${window.location.origin}/${router.currentRoute.value.params.projectNameAbbr}/project-setting/environment/var`, '_blank');
};

const handleFocus = () => {
  store.dispatch('Debug/listServes', {
    serveId: currServe.value.id,
  })
};

watch(
  () => {
    return containerScrollTop && containerScrollTop.value;
  },
  (val) => {
    if (!props.disabled) {
      selectEnvTopPosition.value = getSelectEnvTopPosition();
    }
  }
);

watch(
  () => {
    return props.disabled;
  },
  (val) => {
    if (!val) {
      selectEnvTopPosition.value = getSelectEnvTopPosition();
    }
  }
);
</script>
<style lang="less">
.select-env-container { // related to body
  z-index: 1001;
  right: 22px;
  width: 120px;

  &.fixed {
    position: fixed;
  }

  .ant-select {
    width: 100%;
  }
}
</style>
