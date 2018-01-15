import {Component} from "@angular/core";
import { GlobalState } from '../../global.state';

import { CONSTANT } from '../../utils/constant';
import {RouteService} from "../../service/route";
import {AccountService} from "../../service/account";

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
    {link:'/pages/org-admin/property/case-type/list', title: '属性设置'}
  ];

  menuItems:any[] = this.menus;

  constructor(private _state: GlobalState, private _routeService:RouteService, private accountService: AccountService) {

  }

  ngOnInit() {

  }

}
