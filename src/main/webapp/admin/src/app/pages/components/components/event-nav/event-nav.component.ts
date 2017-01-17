import {Component, ViewEncapsulation, Input, Output, EventEmitter} from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';

@Component({
  selector: 'event-nav',
  encapsulation: ViewEncapsulation.None,
  styles: [require('./event-nav.scss')],
  template: require('./event-nav.html'),
})

export class EventNav {

  @Input() tabModel: any;
  @Input() needCreate: any;

  @Output() createClick = new EventEmitter<any>();
  @Output() itemClick = new EventEmitter<any>();

  constructor(private _router: Router) {
  }

  public onCreateClick($event: any):boolean {
    this.createClick.emit($event);
    return false;
  }
  public onItemClick(tabName: string, $event: any):boolean {
    $event.tabModel = tabName;
    this.itemClick.emit($event);
    return false;
  }
  public back($event: any):void {
    let that = this;
    that._router.navigateByUrl("/pages/event/list");
  }
}
