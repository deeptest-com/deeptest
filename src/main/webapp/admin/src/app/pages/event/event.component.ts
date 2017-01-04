import {Component} from '@angular/core';

@Component({
  selector: 'event',
  styles: [require('./event.scss')],
  template: `<router-outlet></router-outlet>`
})
export class Event {

  constructor() {
  }

  ngOnInit() {
  }

}