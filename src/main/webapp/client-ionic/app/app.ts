import {Component, } from '@angular/core';
import {Platform, ionicBootstrap} from 'ionic-angular';
import {StatusBar} from 'ionic-native';
import {TabsPage} from './pages/tabs/tabs';

import {WebsockService} from './services/websock';
import {Utils} from './utils/utils';

@Component({
  template: '<ion-nav [root]="rootPage"></ion-nav>',
  // config: {mode: 'ios'}
})
export class MyApp {

  private rootPage: any;

  constructor(private platform: Platform) {

    Utils.ClientCofig();
    WebsockService.connect();

    this.rootPage = TabsPage;

    platform.ready().then(() => {
      // Okay, so the platform is ready and our plugins are available.
      // Here you can do any higher level native things you might need.
      StatusBar.styleDefault();
    });
  }
}

ionicBootstrap(MyApp, [], {
  platform: 'ios',
  mode: 'ios',
  backButtonText: '',
  iconMode: 'ios',
  modalEnter: 'modal-slide-in',
  modalLeave: 'modal-slide-out',
  tabsPlacement: 'bottom',
  pageTransition: 'ios',
});
