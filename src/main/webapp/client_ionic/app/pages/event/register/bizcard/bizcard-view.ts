
import {Component} from "@angular/core";
import {NavController} from "ionic-angular";

import {Utils} from '../../../../utils/utils';
import {ImgPathPipe} from "../../../../pipes/img-path";
import {IosDatePipe} from "../../../../pipes/ios-date";
import {CommonService} from "../../../../services/common";
import {PostService} from "../../../../services/post";
import {RegisterService} from "../../../../services/register";
import {BizcardService} from "../../../../services/bizcard";

@Component({
  templateUrl: 'build/pages/event/register/bizcard/bizcard-view.html',
  providers: [CommonService, PostService, RegisterService, BizcardService],
  directives: [],
  pipes: [ImgPathPipe, IosDatePipe]
})
export class BizcardView {
  errorMessage:any;
  data:any;

  constructor(private nav:NavController,
              private _registerService:RegisterService, private _bizcardService:BizcardService, private _commonService:CommonService) {
    let me = this;

  }

  ngOnInit() {
    this.loadData();
  }

  onPageDidEnter():void {

  }

  loadData():void {
    var me = this;
    me._bizcardService.getMyBizcard({eventId: -1}).subscribe((json) => {
      me.data = json.data;
    });
  }

  register():void {
    var me = this;
    me._registerService.register({eventId: -1}).subscribe((json) => {
      me.data = json.data;
    });
  }
    
  backgroundImageStyle(url: string):string {
      return Utils.backgroundImage(url, true);
  }

}
