import {Component} from '@angular/core';
import {NavController} from 'ionic-angular';

import {CONSTANT} from '../../utils/constant';
import {Utils} from '../../utils/utils';
import {ImgPathPipe} from '../../pipes/img-path';
import {ImgSizePipe} from '../../pipes/img-size';

import {PubSubService} from '../../services/pub-sub-service';
import {CommonService}    from '../../services/common';
import {PostService}    from '../../services/post';
import {HomeService}    from '../../services/home';

import {Introduce} from '../event/introduce/introduce';
import {Live} from '../event/live/live';
import {Review} from '../event/review/review';

import {Schedule} from '../event/schedule/schedule';
import {GuestList} from '../event/guest/guest-list';
import {RegisterIndex} from '../event/register/register-index';
import {ServiceList} from '../event/service/service-list';
import {AroundList} from '../event/around/around-list';

// import {SockJS} from 'sockjs-client/lib';

@Component({
  templateUrl: 'build/pages/home/home.html',
  providers: [CommonService, PostService, HomeService],
  directives: [],
  pipes: [ImgPathPipe, ImgSizePipe]
})
export class HomePage {
  errorMessage: any;
  data: any;
  slideHeight: number;
  rowHeight: number;
  mySlideOptions: any

  constructor(private nav: NavController,
              private _homeService: HomeService, private _commonService: CommonService) {
    let me = this;

    me.rowHeight = CONSTANT.W / 3;
    me.slideHeight = CONSTANT.H - 100 - me.rowHeight * 2;

    me.mySlideOptions = {
        pager: false,
        initialSlide: 0,
        autoplay: 1000,
        speed: 1000,
        loop: true
    };
  }

  ngOnInit() {
    this.loadData();
  }

  onPageDidEnter(): void {
    let i = 0;
  }

  loadData():void {
    var me = this;
    this._homeService.getData({eventId: CONSTANT.EVENT_ID}).subscribe((json) => {
          me.data = json;

          CONSTANT.event = json.event;
    });
  }

  gotoSession(): void {
    //this.nav.push(Introduce);
    //this.nav.push(Live); 
    this.nav.push(Review);
  }
  gotoSchedule(): void {
    this.nav.push(Schedule);
  }
  gotoGuest(): void {
    this.nav.push(GuestList);
  }
  gotoRegister(): void {
    this.nav.push(RegisterIndex);
  }
  gotoService(): void {
    this.nav.push(ServiceList);
  }
  gotoAround(): void {
    this.nav.push(AroundList);
  }

  backgroundImageStyle(url: string):string {
      let urlFull = Utils.ImgUrl(Utils.ImgSize(url), true);
    
      return "url(" + urlFull + ")";
  }
}
