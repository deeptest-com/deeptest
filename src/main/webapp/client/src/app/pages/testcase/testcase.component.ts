import {Component} from '@angular/core';

@Component({
  selector: 'testcase',
  styles: [require('./testcase.scss')],
  template: `<router-outlet></router-outlet>`
})
export class Testcase {

  constructor() {
  }

  ngOnInit() {
  }

}