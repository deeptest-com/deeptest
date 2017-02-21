import {Component} from '@angular/core';
import {NavController} from 'ionic-angular';

import {CONSTANT} from '../../../utils/constant';
import {ImgPathPipe} from '../../../pipes/img-path';
import {IosDatePipe} from '../../../pipes/ios-date';

import {PubSubService} from '../../../services/pub-sub-service';
import {CommonService}    from '../../../services/common';
import {PostService}    from '../../../services/post';

import {RegisterService}    from '../../../services/register';
import {SessionService}    from '../../../services/session';

import {BizcardView} from './bizcard/bizcard-view';
import {BizcardEdit} from './bizcard/bizcard-edit';

@Component({
  templateUrl: 'build/pages/event/register/register-index.html',
  providers: [CommonService, PostService, RegisterService, SessionService],
  directives: [],
  pipes: [ImgPathPipe, IosDatePipe]
})
export class RegisterIndex {
  errorMessage: any;
  rowHeight: number;
  data: any;
  selected: any = {};

  constructor(private nav: NavController,
              private _registerService: RegisterService, private _sessionService: SessionService, private _commonService: CommonService) {
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
    me._registerService.getInfo({eventId: -1}).subscribe((json) => {
       me.data = json.data;
    });
  }
    
  goto(page): void {
    var me = this;
      if (page == 'register') {
          this.nav.push(BizcardView);
      } else if(page == 'bizcard') {
          this.nav.push(BizcardEdit);
      }
  }
  
}
