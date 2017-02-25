import {Component} from '@angular/core';

@Component({
  selector: 'business',
  styles: [require('./business.scss')],
  template: `<router-outlet></router-outlet>`
})
export class Business {

  constructor() {
  }

  ngOnInit() {
  }

}
