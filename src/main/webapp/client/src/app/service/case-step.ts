import {Injectable} from "@angular/core";
import {RequestService} from "./request";

@Injectable()
export class CaseStepService {
  constructor(private _reqService: RequestService) {
  }

  _api_url = 'case_step/';

  up(caseStep: any) {
    return this._reqService.post(this._api_url + 'up', caseStep);
  }

  down(caseStep: any) {
    return this._reqService.post(this._api_url + 'down', caseStep);
  }

  save(caseStep: any) {
    return this._reqService.post(this._api_url + 'save', caseStep);
  }

  delete(caseStep: any) {
    return this._reqService.post(this._api_url + 'delete', caseStep);
  }
}



