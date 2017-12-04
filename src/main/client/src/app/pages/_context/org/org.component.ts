import { Component } from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { Routes } from '@angular/router';

import { CONSTANT } from '../../../utils/constant';
import {RouteService} from "../../../service/route";

import { AccountService } from '../../../service/account';
import { ProjectService } from '../../../service/project';

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

  constructor(private _route: ActivatedRoute, private _routeService: RouteService, private _projectService: ProjectService,
              private accountService: AccountService) {

  }

  ngOnInit() {
    this._route.params.subscribe(params => {
      this.orgId = params['orgId'];
    });
    console.log('===Org', this.orgId);

    if (CONSTANT.CURR_ORG_ID != this.orgId) {
      CONSTANT.CURR_ORG_ID = this.orgId;

      this.accountService.loadProfileRemote().subscribe((result: any) => {
        if (!result) {
          this._routeService.navTo('/login');
        }
      });
    }
  }
}
