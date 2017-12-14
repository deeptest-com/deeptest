import {Input, Component} from '@angular/core';

import {BaMsgCenterService} from './baMsgCenter.service';

@Component({
  selector: 'ba-msg-center',
  providers: [BaMsgCenterService],
  styleUrls: ['./baMsgCenter.scss'],
  templateUrl: './baMsgCenter.html'
})
export class BaMsgCenter {

  @Input() alerts;
  @Input() msgs;

  constructor(private _baMsgCenterService:BaMsgCenterService) {

  }

}
