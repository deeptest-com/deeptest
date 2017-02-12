export const PAGES_MENU = [
  {
    path: 'pages',
    children: [
      {
        path: 'dashboard',
        data: {
          menu: {
            title: '首页',
            icon: 'ion-android-home',
            selected: false,
            expanded: false,
            order: 0
          }
        }
      },
      {
        path: 'event',
        data: {
          menu: {
            title: '会议管理',
            icon: 'ion-calendar',
            selected: false,
            expanded: true,
            order: 100,
          }
        },
        children: [
          {
            path: 'list',
            data: {
              menu: {
                title: '我的会议',
              }
            }
          },
          {
            path: 'settings',
            data: {
              menu: {
                title: '全局设置',
              }
            }
          }
        ]
      },
      {
        path: 'business',
        data: {
          menu: {
            title: '业务管理',
            icon: 'ion-settings',
            selected: false,
            expanded: true,
            order: 100,
          }
        },
        children: [
          {
            path: 'company-edit',
            data: {
              menu: {
                title: '公司管理'
              }
            }
          },
          {
            path: 'account-list',
            data: {
              menu: {
                title: '账号管理'
              }
            }
          }
        ]
      },
      {
        path: 'personal',
        data: {
          menu: {
            title: '个人设置',
            icon: 'ion-person',
            selected: false,
            expanded: true,
            order: 100,
          }
        },
        children: [
          {
            path: 'profile',
            data: {
              menu: {
                title: '修改信息'
              }
            }
          },
          {
            path: 'password',
            data: {
              menu: {
                title: '修改密码'
              }
            }
          }
        ]
      }

    ]
  }
];
