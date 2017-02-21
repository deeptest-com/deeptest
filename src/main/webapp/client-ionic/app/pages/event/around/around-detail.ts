import {Component} from '@angular/core';
import {NavController, NavParams} from 'ionic-angular';

import {CONSTANT} from '../../../utils/constant';
import {ImgPathPipe} from '../../../pipes/img-path';
import {IosDatePipe} from '../../../pipes/ios-date';
import {AroundNamePipe} from '../../../pipes/misc';

import {PubSubService} from '../../../services/pub-sub-service';
import {CommonService}    from '../../../services/common';
import {PostService}    from '../../../services/post';

import {MapService}    from '../../../services/map';
import {AroundService}    from '../../../services/around';

@Component({
  templateUrl: 'build/pages/event/around/around-detail.html',
  providers: [PostService, MapService, AroundService],
  directives: [],
  pipes: [ImgPathPipe, IosDatePipe, AroundNamePipe]
})
export class AroundDetail {
  errorMessage: any;
  mapHeight: number;
  rowHeight: number;
  data: any;

  constructor(private nav: NavController, private navParams: NavParams,
              private _arroundService: AroundService, private _mapService: MapService) {
    let me = this;

    me.rowHeight = CONSTANT.W / 3;
    me.data = navParams.get('data');
  }

  ngOnInit() {
    this.loadData();
  }

  onPageDidEnter(): void {

  }

  loadData():void {
    var me = this;



  }

}
