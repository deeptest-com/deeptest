import { Component } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { CONSTANT } from '../../../utils/constant';

@Component({
  selector: 'org',
  template: `
    <div class="org">
      <router-outlet></router-outlet>
    </div>
  `
})
export class Org {
  orgId: number;

  constructor(private _route: ActivatedRoute) {
    // this._route.params.subscribe(params => {
    //   this.orgId = params['orgId'];
    // });
    // console.log('==Current Org', this.orgId);
    //
    // if (CONSTANT.CURR_ORG_ID != this.orgId) {
    //   CONSTANT.CURR_ORG_ID = this.orgId;
    // }
  }
  ngOnInit() {

  }
}
