import {Component} from "@angular/core";
import { GlobalState } from '../../global.state';

@Component({
  selector: 'org-admin',
  styleUrls: ['./org-admin.scss'],
  templateUrl: './org-admin.html'
})
export class OrgAdmin {

  menus:any[] = [
    {link:'/pages/org-admin/org/list', title: '我的组织'},
    {link:'/pages/org-admin/user/list', title: '组织用户'},
    {link:'/pages/org-admin/group/list', title: '组织群组'},
    {link:'/pages/org-admin/org-role/list', title: '组织角色'},
    {link:'/pages/org-admin/project-role/list', title: '项目角色'},
    {link:'/pages/org-admin/property/case-type/list', title: '属性设置'},

    {link2:'/pages/org-admin/settings', title: '站点配置'},
    {link2:'/pages/org-admin/integration', title: '第三方集成'},
    {link2:'/pages/org-admin/license', title: '许可证'}
  ];

  menuItems:any[] = this.menus;

  constructor(private _state: GlobalState) {
    this._state.subscribe('org.ready', (orgReady) => {

      if (!orgReady) {
        this.menuItems = [
          {link:'/pages/org-admin/org/edit/null', title: '新建组织', selected: true}
        ];
      } else {
        this.menuItems = this.menus;
      }
    });
  }

  ngOnInit() {
  }

}
