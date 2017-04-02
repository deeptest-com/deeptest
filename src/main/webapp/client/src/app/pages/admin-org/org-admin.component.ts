import {Component} from "@angular/core";

@Component({
  selector: 'org-admin',
  styles: [require('./org-admin.scss')],
  template: require('./org-admin.html')
})
export class OrgAdmin {

  public menuItems:any[] = [
    {link:'/pages/org-admin/org/list', title: '我的公司'},
    {link:'/pages/org-admin/user/list', title: '公司用户'},
    {link:'/pages/org-admin/group/list', title: '公司群组'},
    {link:'/pages/org-admin/role/list', title: '公司角色'},
    {link:'/pages/org-admin/custom', title: '属性定义'},
    {link:'/pages/org-admin/settings', title: '站点配置'},
    {link:'/pages/org-admin/integration', title: '第三方集成'},
    {link:'/pages/org-admin/license', title: '许可证'}
  ];

  constructor() {
  }

  ngOnInit() {
  }

}
