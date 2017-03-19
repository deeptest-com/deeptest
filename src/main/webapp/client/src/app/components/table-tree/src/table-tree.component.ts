import { Input, Output, EventEmitter, Component, OnInit, ElementRef, Inject } from '@angular/core';

import {GlobalState} from '../../../global.state';

import { CONSTANT } from '../../../utils/constant';

@Component({
  selector: 'table-tree',
  styles: [require('./styles.scss')],
  template: require('./table-tree.html')
})
export class TableTreeComponent implements OnInit {
  @Input()
  public models: any;
  @Input()
  public maxLevel: number;

  @Input()
  public indx: number;
  @Input()
  public childrenNumb: number;


  @Output()
  public action: EventEmitter<any> = new EventEmitter<any>();

  constructor(private _state:GlobalState, private el: ElementRef) {
    let that = this;

  }

  public ngOnInit(): void {
    let nativeElement: HTMLElement = this.el.nativeElement;
    let parentElement: HTMLElement = nativeElement.parentElement;
    // move all children out of the element
    while (nativeElement.firstChild) {
      parentElement.insertBefore(nativeElement.firstChild, nativeElement);
    }
    // remove the empty element(the host)
    parentElement.removeChild(nativeElement);
  }

  public onAction(e: any, act:string, id: Number): void {
      this.action.emit({act:act, id: id});
  }
}
