import {Component} from '@angular/core';
import {NavController} from 'ionic-angular';

import {CONSTANT} from '../../../../utils/constant';
import {ImgPathPipe} from '../../../../pipes/img-path';
import {IosDatePipe} from '../../../../pipes/ios-date';

import {PubSubService} from '../../../../services/pub-sub-service';
import {CommonService}    from '../../../../services/common';
import {PostService}    from '../../../../services/post';
import {ChatService}    from '../../../../services/chat';

@Component({
  templateUrl: 'build/pages/event/live/chat/comments.html',
  providers: [CommonService, PostService, ChatService],
  directives: [],
  pipes: [ImgPathPipe, IosDatePipe]
})
export class Comments {
  errorMessage: any;
  data: any;

  constructor(private nav: NavController,
              private _chatService: ChatService, private _commonService: CommonService) {
    let me = this;

  }

  ngOnInit() {
    this.loadData();
  }

  onPageDidEnter(): void {

  }

  loadData():void {
    var me = this;
//    me._chatService.getData({eventId: -1}).subscribe((json) => {
//          me.data = json.data;
//    });
  }

}
