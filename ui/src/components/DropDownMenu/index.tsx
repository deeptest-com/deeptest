import { PropType, defineComponent, VNode, toRefs, computed } from "vue";
import { useStore } from "vuex";
import { PermissionButtonType } from "@/types/permission";
import "./index.less";
import { MoreOutlined } from "@ant-design/icons-vue";

type Recordable<T = any> = {
  [x: string]: T;
}

interface MenuItem {
  /** 提示 */
  tooltip?: string;
  /** 显示图标，只支持图片 */
  icon?: string;
  /** 操作名称 */
  label?: string | JSX.Element | ((record: Recordable) => VNode | string);
  /** 是否渲染 */
  ifShow?: boolean | ((record: Recordable) => boolean);
  /** 权限编码 */
  auth?: string;
  children?: MenuItem[];
  disabled?: boolean;
  key?: string;
} 

/**
 * props定义
 * 
 */
const DropdownMenuProps = {
  dropdownList: {
    type: Array as PropType<MenuItem[]>,
    default: [],
  }, // 下拉菜单
  actionList: {
    type: Array as PropType<MenuItem[]>,
    default: [],
  }, // 无下拉的菜单
  record: {
    type: Object,
    default: {},
  } // 当前操作项
};

const MenuItem = defineComponent({
  name: 'MenuItem',
  props: {
    auth: {
      type: String,
    },
    label: {
      type: String,
    },
    action: {
      type: Function as PropType<(...args: any[]) => void>
    },
    tip: {
      type: String,
      required: false,
    },
    record: {
      type: Object,
    }
  },
  setup(props, ctx) {
    const store = useStore();
    const permissionButtonMap = computed(() => {
      return store.state.Global.permissionButtonMap;
    });

    const hasPermission = computed(() => {
      if (!props.auth) {
        return true;
      }
      return permissionButtonMap.value[PermissionButtonType[`${props.auth}`]];
    });

    const defaultTip = '暂无权限，请联系管理员';

    const handleClick = e => {
      if (!hasPermission.value) {
        return;
      }
      props.action?.(props.record);
    };

    return () => {
      return (
        <a-menu-item class={{ 'lyapi-drop-menu-item': true, 'has-no-permission': !hasPermission.value }} onClick={e => handleClick(e)}>
          <a-tooltip title={hasPermission.value ? null : (props.tip || defaultTip)} color="#1677ff">
            {props.label}
          </a-tooltip>
        </a-menu-item>
      )
    }
  },
})


const ActionList = (opts: { list: MenuItem[], record: Recordable}) => {
  const { list, record } = opts;
  const handleClick = () => {
    console.log('点击事件', record);
  };
  return (
    <div class="action-list">
      {list.map((action: MenuItem) => {
        <div onClick={() => handleClick()}>{action.label}</div>
      })}
    </div>
  )
};

const DropdownList = defineComponent({
  name: 'DropdownList',
  props: {
    list: {
      type: Array as PropType<MenuItem[]>,
      default: () => [],
    },
    record: {
      type: Object,
      default: () => {},
    }
  },
  setup(props, { slots }) {

    const vslots = {
      default: () => {
        return slots?.default?.() ||  <MoreOutlined />
      },
      overlay: () => {
        return (
          <a-menu>
            {
              props.list.map((e: any, index) => (
                <MenuItem key={index} {...e} record={props.record} />
              ))
            }
          </a-menu>
        )
      }
    }

    return () => {
      return (
        <a-dropdown v-slots={vslots} />
      ) 
    };
  },
})

const ifShow = (actionItem: MenuItem, props) => {
  if (typeof actionItem.ifShow === 'boolean') {
    return actionItem.ifShow;
  } 
  if (typeof actionItem.ifShow === 'function') {
    return actionItem.ifShow(props.record);
  }
  return true;
}

/**
 * dropdownMenu组件
 */
export const DropdownActionMenu = defineComponent({
  name: 'DropdownMenu',
  props: DropdownMenuProps,
  setup(props, { slots }) {
    const { dropdownList, actionList, record } = toRefs(props);

    
    return () => {
      return (
        <div class="drop-down-action-wrap">
          {actionList.value.length > 0 && (
            <ActionList list={actionList.value} record={record.value} />
          )}
          {actionList.value.length > 0 && (
            <a-divider type="vertical" />
          )}
          {dropdownList.value.length > 0 && (
            <DropdownList list={dropdownList.value.filter(e => ifShow(e, props))} record={record.value} v-slots={slots} />
          )}
        </div>
      )
    }
  },
})