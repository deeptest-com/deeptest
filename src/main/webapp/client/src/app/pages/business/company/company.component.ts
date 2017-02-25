import {Component} from '@angular/core';

@Component({
  selector: 'company',
  styles: [require('./company.scss')],
  template: `<router-outlet></router-outlet>`
})
export class Company {

  constructor() {
  }

  ngOnInit() {
  }

}