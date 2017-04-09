import {Component} from "@angular/core";
import { GlobalState } from '../../global.state';

@Component({
  selector: 'org-admin',
  styles: [require('./org-admin.scss')],
  template: require('./org-admin.html')
})
export class OrgAdmin {

  menus:any = {
    '/pages/org-admin/org/list': '我的公司',
    '/pages/org-admin/user/list': '公司用户',
    '/pages/org-admin/group/list': '公司群组',
    '/pages/org-admin/org-role/list': '公司角色',
    '/pages/org-admin/project-role/list': '项目角色',
    '/pages/org-admin/property/case-type/list': '属性设置',

    '/pages/org-admin/settings': '站点配置',
    '/pages/org-admin/integration': '第三方集成',
    '/pages/org-admin/license': '许可证'
  };

  menuItems:any = this.menus;

  constructor(private _state: GlobalState) {
    this._state.subscribe('org.ready', (orgReady) => {

      if (!orgReady) {
        this.menuItems = {link:'/pages/org-admin/org/edit/null', title: '新建公司', selected: true};
      } else {
        this.menuItems = this.menus;
      }
    });
  }

  ngOnInit() {
  }

}
