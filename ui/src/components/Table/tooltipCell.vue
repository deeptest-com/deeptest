<!-- 表格中 自定义cell展示 -->
<script lang="tsx">
/**
 * props: { text, width }
 * 超出 指定width的将会展示tooltip
 */
import { defineComponent, getCurrentInstance, nextTick, onMounted, ref } from "vue";

export default defineComponent({
  props: {
    text: {
      default: "",
    },
    width: {
      default: 0,
    },
  },
  setup(props) {
    const showTooltip = ref(false);
    const textRef = ref();
    const { proxy } :any= getCurrentInstance();
    const setTooltip = () => {
      console.log(1111);
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

    return {
      textRef,
      showTooltip,
      setTooltip,
    };
  },
  render() {
    return (
      <div style={{ width: `${this.width}px`, cursor: this.showTooltip ? 'pointer' : 'unset' }}>
        <a-tooltip placement="top" arrowPointAtCenter={true} title={this.showTooltip ? this.text : null}>
          <div class="out">
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
  line-height: 18px;
}
</style>
