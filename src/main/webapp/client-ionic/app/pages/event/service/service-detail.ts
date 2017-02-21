import {Component} from '@angular/core';
import {NavController, NavParams} from 'ionic-angular';

import {CONSTANT} from '../../../utils/constant';
import {ImgPathPipe} from '../../../pipes/img-path';
import {IosDatePipe} from '../../../pipes/ios-date';
import {ServiceNamePipe} from '../../../pipes/misc';

import {PubSubService} from '../../../services/pub-sub-service';
import {CommonService}    from '../../../services/common';
import {PostService}    from '../../../services/post';

import {ServiceService}    from '../../../services/service';

@Component({
  templateUrl: 'build/pages/event/service/service-detail.html',
  providers: [CommonService, PostService, ServiceService],
  directives: [],
  pipes: [ImgPathPipe, IosDatePipe, ServiceNamePipe]
})
export class ServiceDetail {
  errorMessage: any;
  rowHeight: number;
  data: any;

  constructor(private nav: NavController, private navParams: NavParams,
              private _serviceService: ServiceService, private _commonService: CommonService) {
    let me = this;

    me.rowHeight = CONSTANT.W / 3;
    me.data = navParams.get('data');
  }

  ngOnInit() {

  }

  onPageDidEnter(): void {

  }

}
