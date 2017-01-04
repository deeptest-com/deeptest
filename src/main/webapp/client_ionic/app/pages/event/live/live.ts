import {Component} from '@angular/core';
import {NavController} from 'ionic-angular';

import {CONSTANT} from '../../../utils/constant';
import {ImgPathPipe} from '../../../pipes/img-path';
import {IosDatePipe} from '../../../pipes/ios-date';

import {PubSubService} from '../../../services/pub-sub-service';
import {CommonService}    from '../../../services/common';
import {PostService}    from '../../../services/post';

import {ChatService}    from '../../../services/chat';

import {Wifi}    from './wifi/wifi';
import {Exchange}    from './exchange/exchange';
import {Chat}    from './chat/chat';

@Component({
  templateUrl: 'build/pages/event/live/live.html',
  providers: [CommonService, PostService, ChatService],
  directives: [],
  pipes: [ImgPathPipe, IosDatePipe]
})
export class Live {
  errorMessage: any;
  data: any;
  rowHeight: number;
  newsHeight: number;

  constructor(private nav: NavController,
              private _chatService: ChatService, private _commonService: CommonService) {
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
//    me._chatService.getData({eventId: -1}).subscribe((json) => {
//          me.data = json.data;
//    });
  }
    
  goto(page):void {
    if (page == 'wifi') {
        this.nav.push(Wifi);
    } else if (page == 'exchange') {
        this.nav.push(Exchange);
    }else if (page == 'chat') {
        this.nav.push(Chat);
    }
  }

}
