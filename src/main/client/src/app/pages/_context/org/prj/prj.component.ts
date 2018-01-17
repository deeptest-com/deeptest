import { Component } from '@angular/core';
import { ActivatedRoute } from '@angular/router';

import { CONSTANT } from '../../../../utils/constant';

import { AccountService } from '../../../../service/account';
import { ProjectService } from '../../../../service/project';

@Component({
  selector: 'prj',
  template: `
    <div class="prj">
      <router-outlet></router-outlet>
    </div>
  `
})
export class Prj {
  prjId: number;

  constructor(private _route: ActivatedRoute, private _projectService: ProjectService,
              private accountService: AccountService) {

  }

  ngOnInit() {
    this._route.params.subscribe(params => {
      this.prjId = params['prjId'];
    });
    console.log('==Current Prj', this.prjId, CONSTANT.CURR_PRJ_ID);

    if (CONSTANT.CURR_PRJ_ID != this.prjId) {
      CONSTANT.CURR_PRJ_ID = this.prjId;

      this._projectService.view(CONSTANT.CURR_PRJ_ID).subscribe((json: any) => {

      });
    }
  }
}
