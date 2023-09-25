import { VNode } from "vue";

type Recordable<T = any> = {
  [x: string]: T;
}

export type MenuItem = {
  /** 提示 */
  tooltip?: string;
  /** 显示图标，只支持图片 */
  icon?: string;
  /** 操作名称 */
  label?: string | JSX.Element | ((record: Recordable) => VNode | string);
  /** 是否渲染 */
  ifShow?: boolean | ((record: Recordable, action: MenuItem) => boolean);
  /** 权限编码 */
  auth?: string;
  children?: MenuItem[];
  disabled?: boolean;
  key?: string;
} 