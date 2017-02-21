import {Component} from '@angular/core';
import {NavController} from 'ionic-angular';

import {CONSTANT} from '../../../../utils/constant';
import {ImgPathPipe} from '../../../../pipes/img-path';
import {IosDatePipe} from '../../../../pipes/ios-date';

import {PubSubService} from '../../../../services/pub-sub-service';
import {CommonService}    from '../../../../services/common';
import {PostService}    from '../../../../services/post';

import {FeedbackService}    from '../../../../services/feedback';

@Component({
  templateUrl: 'build/pages/event/review/feedback/feedback.html',
  providers: [CommonService, PostService, FeedbackService],
  directives: [],
  pipes: [ImgPathPipe, IosDatePipe]
})
export class Feedback {
  errorMessage: any;
  data: any;    

  constructor(private nav: NavController,
              private _feedbackService: FeedbackService, private _commonService: CommonService) {
    let me = this;

  }

  ngOnInit() {
    this.loadData();
  }

  onPageDidEnter(): void {

  }

  loadData():void {
    var me = this;
    
//    me._feedbackService.getDetail({eventId: -1}).subscribe((json) => {
//          me.data = json.data;
//    });
  }
  
}
