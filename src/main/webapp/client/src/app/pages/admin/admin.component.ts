import {Component} from "@angular/core";

@Component({
  selector: 'admin',
  styles: [require('./admin.scss')],
  template: require('./admin.html')
})
export class Admin {

  public menuItems:any[] = [
    {link:'/pages/admin/user/list', title: '用户'},
    {link:'/pages/admin/group/list', title: '用户组'},
    {link:'/pages/admin/role/list', title: '角色&权限'},
    {link:'/pages/admin/custom', title: '自定义属性'},
    {link:'/pages/settings', title: '站点配置'},
    {link:'/pages/integration', title: '第三方集成'},
    {link:'/pages/license', title: '许可证'}
  ];

  constructor() {
  }

  ngOnInit() {
  }

}
