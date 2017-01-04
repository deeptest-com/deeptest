import {Component} from '@angular/core';
import {NavController} from 'ionic-angular';

import {CONSTANT} from '../../../utils/constant';
import {ImgPathPipe} from '../../../pipes/img-path';
import {IosDatePipe} from '../../../pipes/ios-date';

import {PubSubService} from '../../../services/pub-sub-service';
import {CommonService}    from '../../../services/common';
import {PostService}    from '../../../services/post';
import {GuestService}    from '../../../services/guest';

import {GuestDetail}    from './guest-detail';

@Component({
  templateUrl: 'build/pages/event/guest/guest-list.html',
  providers: [CommonService, PostService, GuestService],
  directives: [],
  pipes: [ImgPathPipe, IosDatePipe]
})
export class GuestList {
  errorMessage: any;
  data: any;

  constructor(private nav: NavController,
              private _guestService: GuestService, private _commonService: CommonService) {
    let me = this;

  }

  ngOnInit() {
    this.loadData();
  }

  onPageDidEnter(): void {

  }

  loadData():void {
    var me = this;
    me._guestService.list({eventId: -1}).subscribe((json) => {
          me.data = json.guests;
    });
  }
    
  gotoDetail(item):void {
    var me = this;
    me.nav.push(GuestDetail, {item: item});
  }

}
