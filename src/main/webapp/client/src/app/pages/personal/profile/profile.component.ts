import {Component} from '@angular/core';

@Component({
  selector: 'profile',
  styles: [require('./profile.scss')],
  template: `<router-outlet></router-outlet>`
})
export class Profile {

  constructor() {
  }

  ngOnInit() {
  }

}