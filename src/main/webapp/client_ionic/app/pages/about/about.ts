import {Component} from '@angular/core';
import {NavController} from 'ionic-angular';

import {PubSubService} from '../../services/pub-sub-service';
import {CommonService}    from '../../services/common';
import {PostService}    from '../../services/post';
import {AboutService}    from '../../services/about';

@Component({
  templateUrl: 'build/pages/about/about.html'
})
export class AboutPage {
  constructor(private navCtrl: NavController) {
  }
}
