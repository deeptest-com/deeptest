import {Component} from '@angular/core';
import {NavController} from 'ionic-angular';

import {CONSTANT} from '../../../../utils/constant';
import {ImgPathPipe} from '../../../../pipes/img-path';
import {IosDatePipe} from '../../../../pipes/ios-date';

import {PubSubService} from '../../../../services/pub-sub-service';
import {CommonService}    from '../../../../services/common';
import {PostService}    from '../../../../services/post';

import {WifiService}    from '../../../../services/wifi';

@Component({
  templateUrl: 'build/pages/event/live/wifi/wifi.html',
  providers: [CommonService, PostService, WifiService],
  directives: [],
  pipes: [ImgPathPipe, IosDatePipe]
})
export class Wifi {
  errorMessage: any;
  data: any;

  constructor(private nav: NavController,
              private _wifiService: WifiService, private _commonService: CommonService) {
    let me = this;

  }

  ngOnInit() {
    this.loadData();
  }

  onPageDidEnter(): void {

  }

  loadData():void {
    var me = this;
//    me._wifiService.getData({eventId: -1}).subscribe((json) => {
//          me.data = json.data;
//    });
  }

}
