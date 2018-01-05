import {Injectable} from "@angular/core";
import {RequestService} from "./request";

import * as _ from 'lodash';

@Injectable()
export class CaseCommentsService {
  constructor(private _reqService: RequestService) {
  }

  _api_url = 'case_comments/';

  save(caseId:number, summary: string, content: string) {
    return this._reqService.post(this._api_url + 'save', {caseId: caseId, summary: summary, content: content});
  }
  remove(commentsId:number) {
    return this._reqService.post(this._api_url + 'remove', {commentsId: commentsId});
  }

}



