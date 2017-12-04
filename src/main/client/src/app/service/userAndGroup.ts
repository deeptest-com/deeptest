import * as _ from 'lodash';

import {Injectable} from "@angular/core";

import {RequestService} from "./request";

@Injectable()
export class UserAndGroupService {
  constructor(private _reqService: RequestService) { }
  _api_url = 'userAndGroup/';

  search(orgId:number, keywords: string, exceptIds: string[]) {
    let model = {orgId:orgId, keywords: keywords, exceptIds: exceptIds};
    return this._reqService.post(this._api_url + 'search', model);
  }

}

