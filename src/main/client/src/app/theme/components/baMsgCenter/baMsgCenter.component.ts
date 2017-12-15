import {Input, Component} from '@angular/core';

import {CONSTANT} from "../../../utils/constant";
import {RouteService} from "../../../service/route";
import {BaMsgCenterService} from './baMsgCenter.service';

import {RunService} from '../../../service/run';
import {MsgService} from '../../../service/msg';

@Component({
  selector: 'ba-msg-center',
  providers: [BaMsgCenterService],
  styleUrls: ['./baMsgCenter.scss'],
  templateUrl: './baMsgCenter.html'
})
export class BaMsgCenter {

  @Input() alerts;
  @Input() msgs;

  constructor(private _routeService: RouteService, private _baMsgCenterService:BaMsgCenterService,
              private runService:RunService, private msgService:MsgService) {

  }

  readAllAlerts($event) {
    console.log('readAllAlerts');
    this.runService.markAllRead().subscribe((json:any) => {

    });
  }
  readAllMsgs($event) {
    this.msgService.markAllRead().subscribe((json:any) => {

    });
  }
  moreMsgs($event) {
    let url = '/pages/personal/msg/list';
    this._routeService.navTo(url);
  }

}
