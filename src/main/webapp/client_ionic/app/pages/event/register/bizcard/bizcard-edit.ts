import {Component} from '@angular/core';
import {NavController} from 'ionic-angular';

import {CONSTANT} from '../../../../utils/constant';
import {Utils} from '../../../../utils/utils';
import {ImgPathPipe} from '../../../../pipes/img-path';
import {IosDatePipe} from '../../../../pipes/ios-date';

import {PubSubService} from '../../../../services/pub-sub-service';
import {CommonService}    from '../../../../services/common';
import {PostService}    from '../../../../services/post';

import {BizcardService} from "../../../../services/bizcard";
import {RegisterService}    from '../../../../services/register';

@Component({
  templateUrl: 'build/pages/event/register/bizcard/bizcard-edit.html',
  providers: [CommonService, PostService, BizcardService, RegisterService],
  directives: [],
  pipes: [ImgPathPipe, IosDatePipe]
})
export class BizcardEdit {
  errorMessage: any;
  data: any = {};

  constructor(private nav: NavController,
              private _bizcardService: BizcardService, private _registerService: RegisterService, private _commonService: CommonService) {
    let me = this;

  }

  ngOnInit() {
    this.loadData();
  }

  onPageDidEnter(): void {
    
  }

  loadData():void {
    var me = this;
    me._bizcardService.getMyBizcard({eventId: -1}).subscribe((json) => {
      me.data = json.data;
    });
  }

   saveAndRegister():void {
        var me = this;
        me._registerService.register({eventId: -1}).subscribe((json) => {
               me.data = json.data;
        });
    }

    save():void {
        var me = this;
        me._bizcardService.save({eventId: -1}).subscribe((json) => {
               me.data = json.data;
        });
    }

      backgroundImageStyle(url: string):string {
          return Utils.backgroundImage(url, true);
      }
    
}
