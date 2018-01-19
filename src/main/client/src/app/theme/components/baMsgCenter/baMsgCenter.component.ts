import {Input, Component} from '@angular/core';

import {CONSTANT} from "../../../utils/constant";
import {RouteService} from "../../../service/route";
import {BaMsgCenterService} from './baMsgCenter.service';

import {AlertService} from '../../../service/alert';
import {MsgService} from '../../../service/msg';

@Component({
  selector: 'ba-msg-center',
  providers: [BaMsgCenterService],
  styleUrls: ['./baMsgCenter.scss'],
  templateUrl: './baMsgCenter.html'
})
export class BaMsgCenter {

  noReads: number;
  _alerts: any[];
  @Input() set alerts(vals: any[]) {
    this._alerts = vals;
    this.noReads = 0;

    this._alerts.map((item:any) => {
      if (!item.read) {this.noReads += 1};
    });
  };
  @Input() msgs;

  constructor(private _routeService: RouteService, private _baMsgCenterService:BaMsgCenterService,
              private alertService:AlertService, private msgService:MsgService) {

  }

  readAllAlerts($event) {
    let ids: string = '';
    this._alerts.map((item:any) => {
      if (ids != '') {ids += ','}
      ids += item.id;
    });
    this.alertService.markAllRead(ids).subscribe((json:any) => {

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
