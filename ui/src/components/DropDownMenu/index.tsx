import { PropType, defineComponent } from "vue";

/**
 * props定义
 */
const DropdownMenuProps = {
  menuList: {
    type: Array as PropType<any[]>,
  },
  record: {
    type: Object,
  }
};

/**
 * dropdownMenu组件
 */
export const DropdownMenu = defineComponent({
  name: 'DropdownMenu',
  props: DropdownMenuProps,
  setup(props, ctx) {
    return () => {
      return (
        <a-dropdown>
          
        </a-dropdown>
      )
    }
  },
})