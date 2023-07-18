<template>
  <div class="condition-edit-fullscreen">
    <a-modal
        :visible="visible"
        :footer="null"
        @cancel="cancel"
        title="Basic Modal"
        width="100%"
        wrapClassName="dp-full-modal">

      <div class="content">
        <Extractor v-if="model.entityType === ConditionType.extractor"
                   :condition="model"
                   :finish="onCancel"/>

        <Checkpoint v-if="model.entityType === ConditionType.checkpoint"
                    :condition="model"
                    :finish="onCancel" />

        <Script v-if="model.entityType === ConditionType.script"
                :condition="model"
                :finish="onCancel" />
      </div>

    </a-modal>
  </div>
</template>

<script setup lang="ts">
import {computed, defineProps, inject, ref} from "vue";
import {useI18n} from "vue-i18n";

import {ConditionType} from "@/utils/enum";
import Extractor from "./conditions-post/Extractor.vue";
import Checkpoint from "./conditions-post/Checkpoint.vue";
import Script from "./conditions-post/Script.vue";

const props = defineProps({
  visible: {
    required: true,
    type: Boolean,
  },
  model: {
    required: true,
    type: Object,
  },
  onCancel: {
    required: true,
    type: Function,
  },
})

const cancel = () => {
  console.log('cancel')
  props.onCancel()
}

</script>

<style lang="less">
.pre-condition-main {
  .codes {
    height: 100%;
    min-height: 160px;

    .editor {
      height: 100%;
      min-height: 160px;
    }
  }
}
</style>

<style lang="less" scoped>
.pre-condition-main {
  height: 100%;
  display: flex;
  flex-direction: column;

  .head {
    height: 30px;
    padding: 2px 3px;
    border-bottom: 1px solid #d9d9d9;
  }
  .content {
    flex: 1;
    height: calc(100% - 30px);
    overflow-y: auto;

    display: flex;
    &>div {
      height: 100%;
    }

    .codes {
      flex: 1;
    }
    .refer {
      width: 260px;
      padding: 10px;
      overflow-y: auto;

      .title {
        margin-top: 12px;
      }
      .desc {

      }
    }

    .collapse-list {
      height: 100%;
      width: 100%;
      padding: 3px 0;

      .collapse-item {
        width: 100%;
        border: 1px solid #d9d9d9;
        border-bottom: 0;
        border-radius: 2px;

        &:last-child {
          border-radius: 0 0 2px 2px;
          border-bottom: 1px solid #d9d9d9;
        }

        .header {
          height: 38px;
          line-height: 22px;
          padding: 10px;
          background-color: #fafafa;

          display: flex;
          .title {
            flex: 1;
            font-weight: bolder;
          }
          .buttons {
            width: 160px;
            text-align: right;
          }
        }
        .content {
          padding: 16px 10px;
          width: 100%;
        }
      }
    }
  }
}
</style>

<style lang="less" scoped>

</style>
