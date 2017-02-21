import {Component} from '@angular/core';
import {NavController} from 'ionic-angular';

import {CONSTANT} from '../../../utils/constant';
import {ImgPathPipe} from '../../../pipes/img-path';
import {IosDatePipe} from '../../../pipes/ios-date';

import {PubSubService} from '../../../services/pub-sub-service';
import {CommonService}    from '../../../services/common';
import {PostService}    from '../../../services/post';

import {ReviewService}    from '../../../services/review';

import {BizcardList}    from './bizcard/bizcard-list';
import {Document}    from './document/document';
import {Feedback}    from './feedback/feedback';

@Component({
  templateUrl: 'build/pages/event/review/review.html',
  providers: [CommonService, PostService, ReviewService],
  directives: [],
  pipes: [ImgPathPipe, IosDatePipe]
})
export class Review {
  errorMessage: any;
  data: any;
  rowHeight: number;
  newsHeight: number;

  constructor(private nav: NavController,
              private _reviewService: ReviewService, private _commonService: CommonService) {
    let me = this;
      
    me.rowHeight = CONSTANT.W / 3;
    me.newsHeight = CONSTANT.H - 100 - me.rowHeight;
  }

  ngOnInit() {
    this.loadData();
  }

  onPageDidEnter(): void {

  }

  loadData():void {
    var me = this;
    // me._eventService.getDetail({eventId: -1}).subscribe((json) => {
    //       me.data = json.data;
    // });
  }
    
  goto(page):void {
    var me = this;
    
    console.log(page);
    
    if (page == 'bizcard') {
        this.nav.push(BizcardList);
    } else if (page == 'document') {
        this.nav.push(Document);
    }else if (page == 'feedback') {
        this.nav.push(Feedback);
    }
  }

}
