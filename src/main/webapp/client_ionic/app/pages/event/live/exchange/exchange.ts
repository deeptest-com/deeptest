import {Component} from '@angular/core';
import {NavController, ModalController} from 'ionic-angular';
import { FriendComponent } from '../../../../components/modal/friend';

import {CONSTANT} from '../../../../utils/constant';
import {ImgPathPipe} from '../../../../pipes/img-path';
import {IosDatePipe} from '../../../../pipes/ios-date';

import {PubSubService} from '../../../../services/pub-sub-service';
import {CommonService}    from '../../../../services/common';
import {PostService}    from '../../../../services/post';

import {ExchangeService}    from '../../../../services/exchange';

@Component({
  templateUrl: 'build/pages/event/live/exchange/exchange.html',
  providers: [CommonService, PostService, ExchangeService],
  directives: [],
  pipes: [ImgPathPipe, IosDatePipe]
})
export class Exchange {
  errorMessage: any;
  data: any;
  searching: boolean = false;
  searched: boolean = false;
  rowHeight: any;

  constructor(private nav: NavController, private _modalCtrl: ModalController,
              private _exchangeService: ExchangeService, private _commonService: CommonService) {
    let me = this;

    me.rowHeight = CONSTANT.W / 4;
  }

  ngOnInit() {
    this.loadData();
  }

  onPageDidEnter(): void {

  }

  loadData():void {
    var me = this;
//    me._exchangeService.getData({eventId: -1}).subscribe((json) => {
//          me.data = json.data;
//    });
  }

    startSearch(): void {
        console.log('startSearch');
        let me = this;
        me.searched = true;
        me.searching = !me.searching;

        let run = setTimeout(function() {
            me.searching = false;
        }, 5000);
    }

  open(item) {
    console.log(item);
    let modal = this._modalCtrl.create(FriendComponent,
        {},
        {
          showBackdrop: true,
          enableBackdropDismiss: true
        }
    );
    modal.onDidDismiss(data => {
      console.log(data);
    })
    modal.present();
  }

}
