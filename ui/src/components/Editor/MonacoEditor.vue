<template>
  <div class="monaco-editor-vue3" :style="style"></div>
</template>

<script>
import {defineComponent, computed, toRefs, nextTick, ref, onMounted} from 'vue'
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
    interfaceId: Number,
    value: String,
    language: {type: String, default: 'javascript'},
    theme: {type: String, default: 'vs'},
    options: {type: Object, default() {return {};}},

    onExtractor: {type: Function},
    onReplace: {type: Function},
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

      const { interfaceId, value, language, theme, options } = this;
      Object.assign(options, {scrollbar: {
          useShadows: false,
          automaticLayout: true,
          verticalScrollbarSize: 6,
          horizontalScrollbarSize: 6
        }})

      this.editor = monaco.editor[this.diffEditor ? 'createDiffEditor' : 'create'](this.$el, {
        value: value,
        language: language,
        theme: theme,
        ...options
      });

      this.diffEditor && this._setModel(this.value, this.original);

      if (this.options.usedWith === 'response') {
        addExtractAction(this.editor, this.onExtractor)
      } else if (this.options.usedWith === 'request') {
        addReplaceAction(this.editor, this.onReplace)
      }

      // @event `change`
      const editor = this._getEditor()
      // editor.onDidChangeCursorPosition((e) => {
      //   console.log(JSON.stringify(e));
      // });
      // editor.onDidChangeCursorSelection((e) => {
      //   console.log(e);
      //   this.selection = this.editor.getModel().getValueInRange(this.editor.getSelection())
      //   console.log(this.selection)
      // });

      editor.onDidChangeModelContent(event => {
        const value = editor.getValue()

        if (this.value !== value) {
          this.$emit('change', value, event)
        }

        this.formatDocUpdate(editor)

        setTimeout(() => {
          const elems= document.getElementsByClassName('monaco-editor-vue3');
          for(let i=0; i < elems.length; i++) {
            elems[i].style.maxWidth = 0 // elems[i].clientWidth - 200 + 'px'
          }
        }, 100)
      })

      this.$emit('editorDidMount', this.editor)

      setTimeout(() => {
        this.formatDocInit(editor)
      }, 500)
    },

    formatDocInit: (editor) => {
      console.log('format codes - int')
      nextTick(() => {
        editor.getAction('editor.action.formatDocument')?.run()
      })
    },

    formatDocUpdate: debounce((editor) => {
      console.log('format codes - update')
      nextTick(() => {
        editor.getAction('editor.action.formatDocument')?.run()
      })
    }, 1000),

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
      console.log('watch value')
      this.value !== this._getValue() && this._setValue(this.value);
    },

    original() {
      this._setOriginal()
    },

    language() {
      if(!this.editor) return;

      this.formatDocUpdate(this.editor)

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