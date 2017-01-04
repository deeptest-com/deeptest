import {Component} from '@angular/core';
import {NavController} from 'ionic-angular';

import {CONSTANT} from '../../../utils/constant';
import {ImgPathPipe} from '../../../pipes/img-path';
import {IosDatePipe} from '../../../pipes/ios-date';

import {PubSubService} from '../../../services/pub-sub-service';
import {CommonService}    from '../../../services/common';
import {PostService}    from '../../../services/post';

import {EventService}    from '../../../services/event';

@Component({
  templateUrl: 'build/pages/event/introduce/introduce.html',
  providers: [CommonService, PostService, EventService],
  directives: [],
  pipes: [ImgPathPipe, IosDatePipe]
})
export class Introduce {
  errorMessage: any;
  data: any;

  constructor(private nav: NavController,
              private _eventService: EventService, private _commonService: CommonService) {
    let me = this;
    
  }

  ngOnInit() {
    this.loadData();
  }

  onPageDidEnter(): void {

  }

  loadData():void {
    var me = this;
    me._eventService.getDetail({eventId: -1}).subscribe((json) => {
          me.data = json;
    });
  }

}
