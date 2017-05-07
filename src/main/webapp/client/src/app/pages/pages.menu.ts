export const PAGES_MENU = [
  {
    path: 'pages',
    children: [
      {
        path: 'dashboard',
        data: {
          menu: {
            title: 'general.menu.dashboard',
            icon: 'ion-android-home',
            selected: false,
            expanded: false,
            order: 0
          }
        }
      },
      {
        path: 'report',
        data: {
          menu: {
            title: 'general.menu.report',
            icon: 'ion-ios-paper',
            selected: false,
            expanded: false,
            order: 0
          }
        }
      },

      {
        path: 'test-site',
        data: {
          menu: {
            title: 'general.menu.test-site',
            icon: 'ion-wrench',
            selected: false,
            expanded: true,
            order: 10,
          }
        },
        children: [
          {
            path: '',
            data: {
              menu: {
                title: 'general.menu.devsite',
                url: 'http://dev.console.aispeech.com/console/index.html',
                target: '_blank'
              }
            }
          },
          {
            path: '',
            data: {
              menu: {
                title: 'general.menu.testsite',
                url: 'http://test.console.aispeech.com/console/index.html',
                target: '_blank'
              }
            }
          },
          {
            path: '',
            data: {
              menu: {
                title: 'general.menu.demosite',
                url: 'http://demo.console.aispeech.com/console/index.html',
                target: '_blank'
              }
            }
          }
        ]
      },
      {
        path: 'work-site',
        data: {
          menu: {
            title: 'general.menu.work-site',
            icon: 'ion-hammer',
            selected: false,
            expanded: true,
            order: 10,
          }
        },
        children: [
          {
            path: '',
            data: {
              menu: {
                title: 'general.menu.devops',
                url: 'http://172.16.10.9:8000',
                target: '_blank'
              }
            }
          },
          {
            path: '',
            data: {
              menu: {
                title: 'general.menu.gitlab',
                url: 'https://gitlab.spetechcular.com',
                target: '_blank'
              }
            }
          },
          {
            path: '',
            data: {
              menu: {
                title: 'general.menu.jira',
                url: 'https://jira.spetechcular.com',
                target: '_blank'
              }
            }
          },
          {
            path: '',
            data: {
              menu: {
                title: 'general.menu.wiki',
                url: 'https://wiki.spetechcular.com/display/DevPlatform/AISPEECH+DUI',
                target: '_blank'
              }
            }
          }
        ]
      }
    ]
  }
];
