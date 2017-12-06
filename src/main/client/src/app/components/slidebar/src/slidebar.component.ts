import {Component, ElementRef, HostListener, ViewEncapsulation, Input} from '@angular/core';
import {GlobalState} from '../../../global.state';

@Component({
  selector: 'slidebar',
  encapsulation: ViewEncapsulation.None,
  styleUrls: ['./slidebar.scss'],
  templateUrl: './slidebar.html'
})
export class Slidebar {
  @Input()
  menuItems: any;

  constructor(private _state:GlobalState) {

  }

  ngOnInit():void {

  }

  ngAfterViewInit():void {

  }
}
