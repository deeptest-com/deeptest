import {Component} from "@angular/core";

@Component({
  selector: 'sys-admin',
  styleUrls: ['./sys-admin.scss'],
  templateUrl: './sys-admin.html'
})
export class SysAdmin {

  public menuItems:any[] = [
    {link:'/pages/sys-admin/user/list', title: '用户'},
    {link:'/pages/sys-admin/group/list', title: '群组'},
    {link:'/pages/sys-admin/role/list', title: '角色'},
    {link:'/pages/sys-admin/custom', title: '属性定义'},
    {link:'/pages/sys-admin/settings', title: '系统配置'},
    {link:'/pages/sys-admin/integration', title: '第三方集成'}
  ];

  constructor() {
  }

  ngOnInit() {
  }

}
