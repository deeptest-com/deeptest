import {Component} from '@angular/core';
import {NavController} from 'ionic-angular';

import {CONSTANT} from '../../../utils/constant';
import {ImgPathPipe} from '../../../pipes/img-path';
import {IosDatePipe} from '../../../pipes/ios-date';
import {ServiceNamePipe} from '../../../pipes/misc';

import {PubSubService} from '../../../services/pub-sub-service';
import {CommonService}    from '../../../services/common';
import {PostService}    from '../../../services/post';

import {ServiceService}    from '../../../services/service';

import {ServiceDetail}    from './service-detail';

@Component({
  templateUrl: 'build/pages/event/service/service-list.html',
  providers: [CommonService, PostService, ServiceService],
  directives: [],
  pipes: [ImgPathPipe, IosDatePipe, ServiceNamePipe]
})
export class ServiceList {
  errorMessage: any;
  rowHeight: number;
  data: any;

  constructor(private nav: NavController,
              private _serviceService: ServiceService, private _commonService: CommonService) {
    let me = this;

    me.rowHeight = CONSTANT.W / 3;
  }

  ngOnInit() {
    this.loadData();
  }

  onPageDidEnter(): void {

  }

  loadData():void {
    var me = this;
    me._serviceService.getData({eventId: -1}).subscribe((json) => {
          me.data = json.services;
    });
  }

  gotoDetail(type):void {
    var me = this;

    let data = null;
    for (var i=0; i<me.data.length; i++) {
      if (me.data[i]['type'] == type) {
        data = me.data[i];
        break;
      }
    }
    me.nav.push(ServiceDetail, {data: data});
  }

}
