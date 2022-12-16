<template>
  <div class="monaco-editor-vue3" :style="style"></div>
</template>

<script>
import { defineComponent, computed, toRefs } from 'vue'
import * as monaco from 'monaco-editor'
import bus from "@/utils/eventBus";
import settings from "@/config/settings";
import debounce from "lodash.debounce";
import {addExtractAction, addReplaceAction} from "@/components/Editor/service";

export default defineComponent({
  name: "MonacoEditor",
  props: {
    diffEditor: { type: Boolean, default: false },
    width: {type: [String, Number], default: '100%'},
    height: {type: [String, Number], default: '100%'},
    original: String,
    value: String,
    language: {type: String, default: 'javascript'},
    theme: {type: String, default: 'vs'},
    options: {type: Object, default() {return {};}},
  },
  emits: [
    'editorWillMount',
    'editorDidMount',
    'change'
  ],
  setup(props){
    const { width, height } = toRefs(props)

    const style = computed(()=>{
      const fixedWidth = width.value.toString().includes('%') ? width.value : `${width.value}px`
      const fixedHeight = height.value.toString().includes('%')? height.value : `${height.value}px`

      return {
        width: fixedWidth,
        height: fixedHeight,
        'text-align': 'left'
      }
    })
    return {
      style,
    }
  },

  mounted() {
    console.log('editor mounted')

    this.initMonaco()

    const resizeIt = debounce(() => {
      console.log('resizeIt')
      const container = document.getElementsByClassName('response-renderer')[0]
      const size = {width: container.clientWidth, height: container.clientHeight-30}
      // console.log(size)
      this.editor.layout(size)
    }, 500);

    bus.on(settings.eventEditorContainerHeightChanged, () => {
      console.log('resizeIt')
      resizeIt()
    });
  },

  beforeUnmount() {
    console.log('editor beforeUnmount')

    this.editor && this.editor.dispose();
    bus.off(settings.eventEditorContainerHeightChanged)
  },

  methods: {
    initMonaco(){
      this.$emit('editorWillMount', this.monaco)

      const { value, language, theme, options } = this;
      Object.assign(options, {scrollbar: {
          useShadows: false,
          automaticLayout: true,
          verticalScrollbarSize: 6,
          horizontalScrollbarSize: 6
        }})

      console.log('====', options)

      this.editor = monaco.editor[this.diffEditor ? 'createDiffEditor' : 'create'](this.$el, {
        value: value,
        language: language,
        theme: theme,
        ...options
      });

      this.diffEditor && this._setModel(this.value, this.original);

      if (this.options.usedWith === 'response') {
        addExtractAction(this.editor)
      } else if (this.options.usedWith === 'request') {
        addReplaceAction(this.editor)
      }

      // @event `change`
      const editor = this._getEditor()

      editor.onDidChangeModelContent(event => {
        const value = editor.getValue()
        if (this.value !== value) {
          this.$emit('change', value, event)
        }

        setTimeout(() => {
          if (editor.getAction('editor.action.formatDocument'))
            editor.getAction('editor.action.formatDocument').run()
          console.log('format codes')

          const elems= document.getElementsByClassName('monaco-editor-vue3');
          for(let i=0; i < elems.length; i++) {
            elems[i].style.maxWidth = 0 // elems[i].clientWidth - 200 + 'px'
          }
        }, 100)
      })

      this.$emit('editorDidMount', this.editor)

      setTimeout(() => {
        if (editor.getAction('editor.action.formatDocument'))
          editor.getAction('editor.action.formatDocument').run()
        console.log('format codes')
      }, 100)
    },

    _setModel(value, original) {
      const { language } = this;
      const originalModel = monaco.editor.createModel(original, language);
      const modifiedModel = monaco.editor.createModel(value, language);

      this.editor.setModel({
        original: originalModel,
        modified: modifiedModel
      });
    },

    _setValue(value) {
      let editor = this._getEditor();
      if(editor) return editor.setValue(value);
    },

    _getValue() {
      let editor = this._getEditor();
      if(!editor) return '';
      return editor.getValue();
    },

    _getEditor() {
      if(!this.editor) return null;
      return this.diffEditor ? this.editor.modifiedEditor : this.editor;
    },

    _setOriginal(){
      const { original } = this.editor.getModel()
      original.setValue(this.original)
    }
  },

  watch: {
    options: {
      deep: true,
      handler(options) {
        this.editor.updateOptions(options);
      }
    },

    value() {
      this.value !== this._getValue() && this._setValue(this.value);
    },

    original() {
      this._setOriginal()
    },

    language() {
      if(!this.editor) return;

      if (this.diffEditor) {
        const { original, modified } = this.editor.getModel();
        monaco.editor.setModelLanguage(original, this.language);
        monaco.editor.setModelLanguage(modified, this.language);
      } else
        monaco.editor.setModelLanguage(this.editor.getModel(), this.language);
    },

    theme() {
      monaco.editor.setTheme(this.theme);
    },
  }
});
</script>

<style lang="less">
.monaco-editor-vue3 {
  .monaco-editor {
  }
}
</style>