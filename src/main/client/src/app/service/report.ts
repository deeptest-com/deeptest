import * as _ from 'lodash';

import {Injectable} from "@angular/core";

import {RequestService} from "./request";

@Injectable()
export class ReportService {
  constructor(private _reqService: RequestService) { }
  _api_url = 'report/';

  projectReport(projectId: number) {
    return this._reqService.post(this._api_url + 'project', {projectId: projectId});
  }

  planReport(planId: number) {
    return this._reqService.post(this._api_url + 'plan', {planId: planId});
  }
}
