<!-- 表格中 自定义cell展示 -->
<script lang="tsx">
/**
 * props: { text, width }
 * 超出 指定width的将会展示tooltip
 */
import { defineComponent, getCurrentInstance, nextTick, onMounted, ref, watch } from "vue";

export default defineComponent({
  props: {
    text: {
      default: "",
    },
    width: {
      default: 0,
      required: false,
    },
    maxWidth: {
      default: 0,
      required: false,
    },
    tip: {
      default: "",
      required: false,
    },
    customClass: {
      default: "",
      required: false,
      type: String,
    }
  },
  setup(props) {
    const showTooltip = ref(false);
    const textRef = ref();
    const { proxy } :any= getCurrentInstance();
    const setTooltip = () => {
      nextTick(() => {
        const outElWidth = proxy.$el && proxy.$el.offsetWidth;
        const textElWidth = textRef.value && textRef.value.offsetWidth;

        if (outElWidth < textElWidth) {
          showTooltip.value = true;
        } else {
          showTooltip.value = false;
        }
      });
    };

    onMounted(() => {
      setTooltip();
    })

    watch(() => props.text, () => {
      setTooltip();
    })

    return {
      textRef,
      showTooltip,
      setTooltip,
    };
  },
  render() {
    return (
      <div style={{ width: this.width ? `${this.width}px` : 'max-content', maxWidth: (this.maxWidth || this.width) ?`${this.maxWidth || this.width}px` : '100%' ,cursor: this.showTooltip ? 'pointer' : 'unset' }}>
        <a-tooltip placement="top" arrowPointAtCenter={true} title={this.showTooltip ? this.tip || this.text : null}>
          <div class={['out', this.customClass]}>
            <span ref="textRef" class="text">
              {this.text}
            </span>
          </div>
        </a-tooltip>
      </div>
    );
  },
});
</script>
<style scoped lang="less">
.out {
  word-break: break-word;
  overflow: hidden;
  white-space: nowrap;
  text-overflow: ellipsis;
  height: 100%;

  &.processor_logic_if {
    color: #52c41a;
  }

  &.processor_logic_else {
    color: #f5222d;
  }

  &.disabled {
    color: rgba(0, 0, 0, 0.25) !important;
  }
}
</style>
