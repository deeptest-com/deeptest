import { Input, Output, EventEmitter, Component, OnInit, ElementRef, Inject } from '@angular/core';

import {GlobalState} from '../../../global.state';

import { CONSTANT } from '../../../utils/constant';

@Component({
  selector: '[table-row]',
  template: require('./table-row.html')
})
export class TableRowComponent implements OnInit {
  @Input()
  public model: any;
  @Input()
  public maxLevel: number;

  constructor(private _state:GlobalState, private el: ElementRef) {
    let that = this;

  }

  public ngOnInit(): void {
    // let nativeElement: HTMLElement = this.el.nativeElement;
    // let parentElement: HTMLElement = nativeElement.parentElement;
    // // move all children out of the element
    // while (nativeElement.firstChild) {
    //   parentElement.insertBefore(nativeElement.firstChild, nativeElement);
    // }
    // // remove the empty element(the host)
    // parentElement.removeChild(nativeElement);
  }

}
