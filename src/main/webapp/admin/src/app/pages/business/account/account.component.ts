import {Component} from '@angular/core';

@Component({
  selector: 'account',
  styles: [require('./account.scss')],
  template: `<router-outlet></router-outlet>`
})
export class Account {

  constructor() {
  }

  ngOnInit() {
  }

}