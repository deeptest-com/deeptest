import {Injectable} from "@angular/core";
import {RequestService} from "./request";

import * as _ from 'lodash';

@Injectable()
export class CaseCommentsService {
  constructor(private _reqService: RequestService) {
  }

  _api_url = 'case_comments/';

  save(caseId:number, comment: string) {
    _.merge(comment, {testCaseId: caseId});
    return this._reqService.post(this._api_url + 'save', comment);
  }
  remove(id:number) {
    return this._reqService.post(this._api_url + 'delete', {id: id});
  }

}



