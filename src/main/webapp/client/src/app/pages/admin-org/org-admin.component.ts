import {Component} from "@angular/core";
import { GlobalState } from '../../global.state';

@Component({
  selector: 'org-admin',
  styles: [require('./org-admin.scss')],
  template: require('./org-admin.html')
})
export class OrgAdmin {

  menus:any[] = [
    {link:'/pages/org-admin/org/list', title: '我的公司', selected: true},
    {link:'/pages/org-admin/user/list', title: '公司用户'},
    {link:'/pages/org-admin/group/list', title: '公司群组'},
    {link:'/pages/org-admin/org-role/list', title: '公司角色'},
    {link:'/pages/org-admin/project-role/list', title: '项目角色'},
    {link:'/pages/org-admin/custom', title: '属性定义'},
    {link:'/pages/org-admin/settings', title: '站点配置'},
    {link:'/pages/org-admin/integration', title: '第三方集成'},
    {link:'/pages/org-admin/license', title: '许可证'}
  ];

  menuItems:any[] = this.menus;

  constructor(private _state: GlobalState) {
    this._state.subscribe('org.ready', (orgReady) => {

      if (!orgReady) {
        this.menuItems = [
          {link:'/pages/org-admin/org/edit/null', title: '新建公司', selected: true}
        ];
      } else {
        this.menuItems = this.menus;
      }
    });
  }

  ngOnInit() {
  }

}
