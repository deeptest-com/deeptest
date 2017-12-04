import {Injectable} from "@angular/core";
import {RequestService} from "./request";

import * as _ from 'lodash';

@Injectable()
export class CaseStepService {
  constructor(private _reqService: RequestService) {
  }

  _api_url = 'case_step/';

  up(json: any) {
    return this._reqService.post(this._api_url + 'up', json);
  }

  down(json: any) {
    return this._reqService.post(this._api_url + 'down', json);
  }

  save(testCaseId:number, caseStep: any) {
    _.merge(caseStep, {testCaseId: testCaseId});
    return this._reqService.post(this._api_url + 'save', caseStep);
  }

  delete(caseStep: any) {
    return this._reqService.post(this._api_url + 'delete', caseStep);
  }
}



