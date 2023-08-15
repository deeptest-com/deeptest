<template>
  <a-modal
      :title="'配置' + t(model.entityType || 'empty')"
      :visible="visible"
      :footer="null"
      @cancel="cancel"
      width="100%"
      wrapClassName="dp-full-modal processor-edit-fullscreen">

    <div class="content">
      <div class="processor-form">
        <ProcessorCustomCodeEdit 
          v-if="model.processorType === ProcessorCategory.ProcessorCustomCode"
          :processor="model"
          @cancel="cancel"
        />
      </div>
    </div>

  </a-modal>
</template>

<script setup lang="ts">
import {computed, defineProps, inject, ref} from "vue";
import {ProcessorCategory} from "@/utils/enum";
import ProcessorCustomCodeEdit from "../proccessors/custom_code/edit.vue";
import bus from "@/utils/eventBus";
import settings from "@/config/settings";

import {useI18n} from "vue-i18n";
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
})

const emits = defineEmits(['updateScreen']);

const save = (item) => {
  console.log('save', item)
  bus.emit(settings.eventConditionSave, {});
}

const cancel = () => {
  emits('updateScreen', false);
}

const wrapperCol = { span: 18, offset:4 }

</script>

<style lang="less">
.processor-edit-fullscreen {
  height: 100%;

  .head {
    height: 30px;
    padding: 2px 3px;
    border-bottom: 1px solid #d9d9d9;
  }
  .content {
    height: calc(100% - 30px);

    .processor-form {
      height: calc(100% - 36px);
    }
    .buttons {
      height: 36px;
    }
  }
}
</style>

