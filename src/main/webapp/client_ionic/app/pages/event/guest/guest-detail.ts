import {Component} from '@angular/core';
import {NavController, NavParams} from 'ionic-angular';

import {CONSTANT} from '../../../utils/constant';
import {Utils} from '../../../utils/utils';
import {ImgPathPipe} from '../../../pipes/img-path';
import {IosDatePipe} from '../../../pipes/ios-date';

import {PubSubService} from '../../../services/pub-sub-service';
import {CommonService}    from '../../../services/common';
import {PostService}    from '../../../services/post';
import {GuestService}    from '../../../services/guest';

@Component({
  templateUrl: 'build/pages/event/guest/guest-detail.html',
  providers: [CommonService, PostService, GuestService],
  directives: [],
  pipes: [ImgPathPipe, IosDatePipe]
})
export class GuestDetail {
  errorMessage: any;
  item: any;

  constructor(private nav: NavController, private navParams: NavParams,
              private _guestService: GuestService, private _commonService: CommonService) {
    let me = this;
    
    me.item = navParams.get('item');
    console.log(me.item);
  }

  ngOnInit() {
    this.loadData();
  }

  onPageDidEnter(): void {

  }

  loadData():void {
    var me = this;
  }

  backgroundImageStyle(url: string):string {
      let urlFull = Utils.ImgUrl(Utils.ImgSize(url), true);
        
      return "url(" + urlFull + ")";
  }
}
