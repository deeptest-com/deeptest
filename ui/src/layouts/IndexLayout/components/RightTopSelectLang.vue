<template>
  <div>
    <a-modal title="请选择语言"
             :visible="isVisible"
             :onCancel="onClose">
      <div>
        <a-radio-group name="radioGroup" v-model:value="locale">
          <a-radio v-for="item in locales" @change="changeLang(item)" :key="item" :value="item" :style="radioStyle">
            <!-- <span role="img" :aria-label="languageLabels[item]">{{ languageIcons[item] }}</span> -->
            <span style="margin-left: 5px;">{{ languageLabels[item] }}</span>
          </a-radio>
        </a-radio-group>
      </div>

      <template #footer>
        <a-button @click="onClose" type="primary">完成</a-button>
      </template>
    </a-modal>
  </div>
</template>
<script lang="ts">
import {defineComponent, WritableComputedRef, Ref, ref, PropType} from "vue";
import {setI18nLanguage} from "@/config/i18n";
import {useI18n} from "vue-i18n";

export default defineComponent({
  name: 'SelectLang',
  components: {},
  props: {
    isVisible: {
      type: Boolean,
      required: true
    },
    onClose: {
      type: Function,
      required: true,
    },
  },

  setup(props) {
    const {locale} = useI18n();

    const radioStyle = ref({
      display: 'block',
      height: '30px',
      lineHeight: '30px',
    })

    const locales: string[] = ['zh-CN', 'en-US'];
    const languageLabels: { [key: string]: string } = {
      'zh-CN': '简体中文',
      'en-US': 'English',
    };
    const languageIcons: { [key: string]: string } = {
      'zh-CN': '🇨🇳',
      'en-US': '🇺🇸',
    };

    const changeLang = (key): void => {
      console.log(key)
      setI18nLanguage(key);
    }

    return {
      locales,
      languageLabels,
      languageIcons,
      changeLang,
      locale,
      radioStyle,

    }
  }
})
</script>
<style lang="less" scoped>
.menu {
  .anticon {
    margin-right: 8px;
  }

  .ant-dropdown-menu-item {
    min-width: 160px;
  }
}

.dropDown {
  display: inline-block;
  width: 90px;
  cursor: pointer;
  font-size: 14px;
}
</style>