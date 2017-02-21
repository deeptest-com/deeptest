import {Component} from '@angular/core';
import {NavController} from 'ionic-angular';

import {CONSTANT} from '../../../utils/constant';
import {ImgPathPipe} from '../../../pipes/img-path';
import {IosDatePipe} from '../../../pipes/ios-date';
import {AroundNamePipe} from '../../../pipes/misc';

import {PubSubService} from '../../../services/pub-sub-service';
import {CommonService}    from '../../../services/common';
import {PostService}    from '../../../services/post';

import {MapService}    from '../../../services/map';
import {AroundService}    from '../../../services/around';

import {AroundDetail}    from './around-detail';

@Component({
  templateUrl: 'build/pages/event/around/around-list.html',
  providers: [PostService, MapService, AroundService],
  directives: [],
  pipes: [ImgPathPipe, IosDatePipe, AroundNamePipe]
})
export class AroundList {
  errorMessage: any;
  mapHeight: number;
  rowHeight: number;
  data: any;

  constructor(private nav: NavController,
              private _aroundService: AroundService, private _mapService: MapService) {
    let me = this;

    me.rowHeight = CONSTANT.W / 3;
    me.mapHeight = CONSTANT.H - 100 - me.rowHeight * 2; // w / 1.7;
  }

  ngOnInit() {
    this.loadData();
  }

  onPageDidEnter(): void {

  }

  loadData():void {
    let me = this;
    this._mapService.loadMap(CONSTANT.event.city, CONSTANT.event.place);

    me._aroundService.getData({eventId: -1}).subscribe((json) => {
      me.data = json.data;
    });
  }

  gotoDetail(type):void {
    var me = this;

    let data = null;
    for (var i = 0; i < me.data.length; i++) {
      if (me.data[i]['type'] == type) {
        data = me.data[i];
        break;
      }
    }
    me.nav.push(AroundDetail, {data: data});
    console.log(me.data);
  }

}
