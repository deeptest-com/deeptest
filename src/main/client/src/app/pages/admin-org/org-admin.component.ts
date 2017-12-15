import {Component} from "@angular/core";
import { GlobalState } from '../../global.state';

@Component({
  selector: 'org-admin',
  styleUrls: ['./org-admin.scss'],
  templateUrl: './org-admin.html'
})
export class OrgAdmin {

  menus:any[] = [
    {link:'/pages/org-admin/org/listByPage', title: '我的组织'},
    {link:'/pages/org-admin/user/listByPage', title: '组织用户'},
    {link:'/pages/org-admin/group/listByPage', title: '组织群组'},
    {link:'/pages/org-admin/org-role/listByPage', title: '组织角色'},
    {link:'/pages/org-admin/project-role/listByPage', title: '项目角色'},
    {link:'/pages/org-admin/property/case-type/listByPage', title: '属性设置'}
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
