import {Component} from '@angular/core';
import {HomePage} from '../home/home';
import {FindPage} from '../find/find';
import {AboutPage} from '../about/about';

@Component({
  templateUrl: 'build/pages/tabs/tabs.html'
})
export class TabsPage {

  private homeRoot: any;
  private findRoot: any;
  private aboutRoot: any;

  constructor() {
    // this tells the tabs component which Pages
    // should be each tab's root Page
    this.homeRoot = HomePage;
    this.findRoot = FindPage;
    this.aboutRoot = AboutPage;
  }
}
