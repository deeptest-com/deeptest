<template>
  <div class="monaco-editor-vue3" :style="style"></div>
</template>

<script>
import {defineComponent, computed, toRefs, nextTick, inject} from 'vue'
import * as monaco from 'monaco-editor'
import bus from "@/utils/eventBus";
import settings from "@/config/settings";
import debounce from "lodash.debounce";
import {addExtractAction, addReplaceAction} from "@/components/Editor/service";
import {getSnippet} from "@/views/component/debug/service";

import {UsedBy} from "@/utils/enum";

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
    readOnly:  { type: Boolean, default: false ,required: false},
    onExtractor: {type: Function},
    onReplace: {type: Function},
    hooks: String,
    timestamp: String,
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

    bus.on(settings.eventEditorAction, (data) => {
      console.log('eventEditorAction', data)
      if (data.act === settings.eventTypeContainerHeightChanged) {
        resizeIt()
      } else if (data.act === settings.eventTypeFormat) {
        this.formatDocUpdate(this.editor)
      }
    });
  },

  beforeUnmount() {
    console.log('editor beforeUnmount')

    this.editor && this.editor.dispose();
    bus.off(settings.eventEditorAction)
  },
  unmounted() {
    console.log('editor unmounted')
    this.editor && this.editor.dispose();
    bus.off(settings.eventEditorAction)
  },

  methods: {
    initMonaco() {
      console.log('initMonaco ...')
      this.$emit('editorWillMount', this.monaco)

      const {value, language, theme, options} = this;
      Object.assign(options, {
        scrollbar: {
          useShadows: false,
          automaticLayout: true,
          verticalScrollbarSize: 6,
          horizontalScrollbarSize: 6
        }
      })

      this.editor = monaco.editor[this.diffEditor ? 'createDiffEditor' : 'create'](this.$el, {
        value: value,
        language: language,
        theme: theme,
        ...options,
      });

      this.diffEditor && this._setModel(this.value, this.original);


      // const usedBy = inject('usedBy')
      // if (usedBy === UsedBy.InterfaceDebug) {
      if (this.options.usedWith === 'response') {
        addExtractAction(this.editor, this.onExtractor)
      } else if (this.options.usedWith === 'request') {
        addReplaceAction(this.editor, this.onReplace)
      }
      // }

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
          // 添加最后最后一个参数，标识是否有语法错误
          this.$emit('change', value, event, monaco?.editor?.getModelMarkers({})?.length === 0)
        }

        this.formatDocUpdate(editor)

        setTimeout(() => {
          const elems = document.getElementsByClassName('monaco-editor-vue3');
          for (let i = 0; i < elems.length; i++) {
            elems[i].style.maxWidth = 0 // elems[i].clientWidth - 200 + 'px'
          }
        }, 100)
      })

      this.$emit('editorDidMount', this.editor)

      setTimeout(() => {
        this.formatDocInit(editor)
      }, 500)

      if (options.initTsModules) {
        getSnippet('global').then(json => {
          if (json.code === 0) {
            // const externalDtsFileName = 'ex.d.ts';
            // monaco.languages.typescript.typescriptDefaults.addExtraLib(libSource, `inmemory://modelRef/${externalDtsFileName}`);
            monaco.languages.typescript.typescriptDefaults.addExtraLib(json.data.script);
          }
        })
      }
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
    hooks: {
      handler(val) {
        alert(val)
      }
    },

    options: {
      deep: true,
      handler(options) {
        if (!this.editor) return
        this.editor.updateOptions(options);
      }
    },

    timestamp() {
      console.log('watch timestamp', this.value);
      this.value !== this._getValue() && this._setValue(this.value);
    },
    // value() {
    //   console.log('watch value', this.value);
    //   this.value !== this._getValue() && this._setValue(this.value);
    // },

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

