import {Component} from '@angular/core';

@Component({
  selector: 'personal',
  styles: [require('./personal.scss')],
  template: `<router-outlet></router-outlet>`
})
export class Personal {

  constructor() {
  }

  ngOnInit() {
  }

}
