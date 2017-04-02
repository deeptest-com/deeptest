import {Component, ElementRef, HostListener, ViewEncapsulation, Input} from '@angular/core';
import {GlobalState} from '../../../global.state';
import {layoutSizes} from '../../../theme';

@Component({
  selector: 'slidebar',
  encapsulation: ViewEncapsulation.None,
  styles: [require('./slidebar.scss')],
  template: require('./slidebar.html')
})
export class Slidebar {
  @Input()
  public menuItems: any[];

  @Input()
  companies: any[];
  company: any;

  constructor(private _elementRef:ElementRef, private _state:GlobalState) {

  }

  public ngOnInit():void {

  }

  public ngAfterViewInit():void {

  }

  companyChange():void {

  }
}
