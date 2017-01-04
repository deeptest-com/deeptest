import {Component} from '@angular/core';
import {NavController} from 'ionic-angular';

import {CONSTANT} from '../../../utils/constant';
import {ImgPathPipe} from '../../../pipes/img-path';
import {IosDatePipe} from '../../../pipes/ios-date';

import {TimelineComponent} from '../../../components/timeline/timeline';

import {PubSubService} from '../../../services/pub-sub-service';
import {CommonService}    from '../../../services/common';
import {PostService}    from '../../../services/post';

import {ScheduleService}    from '../../../services/schedule';

@Component({
  templateUrl: 'build/pages/event/schedule/schedule.html',
  providers: [CommonService, PostService, ScheduleService],
  directives: [TimelineComponent],
  pipes: [ImgPathPipe, IosDatePipe]
})
export class Schedule {
  errorMessage: any;
  json: any;
  data: any;
  mode: string = 'bySession';

  constructor(private nav: NavController,
              private _scheduleService: ScheduleService, private _commonService: CommonService) {
    let me = this;
    
  }

  ngOnInit() {
    this.loadData();
  }

  onPageDidEnter(): void {

  }

  onItemSelected(item) {
    // this.nav.push(ProductDetail, item.id);
  }

  loadData() {
    let me = this;

    me._scheduleService.getData({eventId: -1}).subscribe((json) => {
       me.json = json;
       me.show(me.mode);
    });
  }
    
  show(by) {
    let me = this;
      
    me.mode = by;
    me.data = me.json[me.mode]
  }

}
