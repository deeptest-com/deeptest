<template>
  <a-modal
      :title="'编辑Mock脚本'"
      :visible="visible"
      :footer="null"
      @cancel="cancel"
      width="100%"
      wrapClassName="dp-full-modal mock-script-fullscreen"
      :bodyStyle="{height: 'calc(100% - 36px)'}">

    <div class="content">
      <div class="form">
        <MockScript />
      </div>

      <div class="buttons">
        <a-form-item :wrapper-col="wrapperCol">
          <a-button type="primary" @click="save">保存</a-button>
          <a-button @click="cancel" style="margin-left:10px">取消</a-button>
        </a-form-item>
      </div>
    </div>

  </a-modal>
</template>

<script setup lang="ts">
import {computed, defineProps, inject, provide, ref} from "vue";
import {useI18n} from "vue-i18n";

import MockScript from "./Script.vue";
import bus from "@/utils/eventBus";
import settings from "@/config/settings";
import {UsedBy} from "@/utils/enum";
const {t} = useI18n();

const props = defineProps({
  visible: {
    required: true,
    type: Boolean,
  },
  model: {
    required: true,
    type: Object,
  },
  onFinish: {
    required: true,
    type: Function,
  },
  onCancel: {
    required: true,
    type: Function,
  },
})

const save = () => {
  console.log('save')
  props.onFinish()
}

const cancel = () => {
  console.log('cancel')
  props.onCancel()
}

const wrapperCol = { span: 18, offset:4 }

</script>

<style lang="less">
.mock-script-fullscreen {
  height: 100%;

  .head {
    height: 30px;
    padding: 2px 3px;
    border-bottom: 1px solid #d9d9d9;
  }
  .content {
    height: calc(100% - 30px);

    .form {
      height: calc(100% - 36px);
    }
    .buttons {
      height: 36px;
    }
  }
}
</style>

