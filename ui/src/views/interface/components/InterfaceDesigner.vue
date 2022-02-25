<template>
  <div class="main">
    {{ model.name }}
  </div>
</template>

<script lang="ts">
import {defineComponent, PropType, Ref} from "vue";
import {useI18n} from "vue-i18n";
import {Form, message} from 'ant-design-vue';
import {CloseOutlined, PlusOutlined} from "@ant-design/icons-vue";

const useForm = Form.useForm;

interface InterfaceDesignerSetupData {
  submit: (e) => void;
}

export default defineComponent({
  name: 'InterfaceDesigner',
  props: {
    model: {
      type: Object,
      required: true
    },
    onSubmit: {
      type: Function as PropType<(model: any) => void>,
      required: true
    }
  },
  components: {
  },
  setup(props): InterfaceDesignerSetupData {
    const {t} = useI18n();

    const submit = (e) => {
      console.log('submit')

      props.onSubmit(props.model);
    };

    return {
      submit
    }
  }
})
</script>

<style lang="less" scoped>
.main {
  padding: 16px;
}
</style>