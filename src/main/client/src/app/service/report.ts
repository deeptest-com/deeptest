import * as _ from 'lodash';

import {Injectable} from "@angular/core";

import {RequestService} from "./request";

@Injectable()
export class ReportService {
  constructor(private _reqService: RequestService) { }
  _api_url = 'report/';

  list(query:any, pageNumb: number, pageSize: number) {
    _.merge(query, {pageNumb: pageNumb, pageSize: pageSize});
    return this._reqService.post(this._api_url + 'list', query);
  }

  get(id: number) {
    let model = {id: id};
    return this._reqService.post(this._api_url + 'get', model);
  }
}
