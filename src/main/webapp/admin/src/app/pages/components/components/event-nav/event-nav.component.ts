import {Component, ViewEncapsulation, Input, Output, EventEmitter} from '@angular/core';

@Component({
  selector: 'event-nav',
  encapsulation: ViewEncapsulation.None,
  styles: [require('./event-nav.scss')],
  template: require('./event-nav.html'),
})

export class EventNav {

  @Input() tabModel: any;

  @Output() itemClick = new EventEmitter<any>();

  constructor() {
  }

  public onItemClick(tabName: string, $event: any):boolean {
    $event.tabModel = tabName;
    this.itemClick.emit($event);
    return false;
  }

}
