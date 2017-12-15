import {Component} from "@angular/core";

@Component({
  selector: 'sys-admin',
  styleUrls: ['./sys-admin.scss'],
  templateUrl: './sys-admin.html'
})
export class SysAdmin {

  public menuItems:any[] = [
    {link2:'/pages/sys-admin/user/list', title: '系统用户'},
    {link2:'/pages/sys-admin/group/list', title: '系统群组'},
    {link2:'/pages/sys-admin/role/list', title: '系统角色'},

    {link2:'/pages/sys-admin/integration', title: '第三方集成'},
    {link2:'/pages/sys-admin/settings', title: '系统配置'},
    {link2:'/pages/sys-admin/license', title: '许可证'}
  ];

  constructor() {
  }

  ngOnInit() {
  }

}
