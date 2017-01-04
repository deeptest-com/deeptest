import {Component} from '@angular/core';
import {NavController} from 'ionic-angular';

import {CONSTANT} from '../../../../utils/constant';
import {ImgPathPipe} from '../../../../pipes/img-path';
import {IosDatePipe} from '../../../../pipes/ios-date';

import {PubSubService} from '../../../../services/pub-sub-service';
import {CommonService}    from '../../../../services/common';
import {PostService}    from '../../../../services/post';

import {BizcardService}    from '../../../../services/bizcard';

@Component({
  templateUrl: 'build/pages/event/review/bizcard/bizcard-list.html',
  providers: [CommonService, PostService, BizcardService],
  directives: [],
  pipes: [ImgPathPipe, IosDatePipe]
})
export class BizcardList {
  errorMessage: any;
  data: any;
  rowHeight: any;

  constructor(private nav: NavController,
              private _bizcardService: BizcardService, private _commonService: CommonService) {
    let me = this;

    me.rowHeight = CONSTANT.W / 4;
  }

  ngOnInit() {
    this.loadData();
  }

  onPageDidEnter(): void {

  }

  loadData():void {
    var me = this;
//    me._bizcardService.getDetail({eventId: -1}).subscribe((json) => {
//          me.data = json.data;
//    });
  }

}
