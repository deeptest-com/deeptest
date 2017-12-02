import {Component} from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { CONSTANT } from '../../../utils/constant';
import { Utils } from '../../../utils/utils';

import { AccountService } from '../../../service/account';
import { ProjectService } from '../../../service/project';

@Component({
  selector: 'plan',
  styleUrls: ['./plan.scss'],
  templateUrl: './plan.html'
})
export class Plan {
  projectId: number;

  constructor(private _route: ActivatedRoute, private _projectService: ProjectService, private accountService: AccountService) {
    this._route.params.forEach((params: Params) => {
      this.projectId = +params['projectId'];
    });

    this._projectService.view(this.projectId).subscribe((json:any) => {
      this.accountService.changeRecentProjects(json.recentProjects);
    });
  }

  ngOnInit() {

  }

}
