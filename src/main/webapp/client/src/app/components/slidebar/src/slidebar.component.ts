import {Component, ElementRef, HostListener, ViewEncapsulation, Input} from '@angular/core';
import {GlobalState} from '../../../global.state';
import {layoutSizes} from '../../../theme';

@Component({
  selector: 'slidebar',
  encapsulation: ViewEncapsulation.None,
  styleUrls: ['./slidebar.scss'],
  templateUrl: './slidebar.html'
})
export class Slidebar {
  @Input()
  public menuItems: any;

  @Input()
  companies: any[];
  company: any;
  orgReady: boolean = false;

  constructor(private _state:GlobalState) {

  }

  public ngOnInit():void {

  }

  public ngAfterViewInit():void {

  }

  companyChange():void {

  }
}
