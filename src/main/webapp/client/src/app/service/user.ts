import * as _ from 'lodash';

import {Injectable} from "@angular/core";

import {RequestService} from "./request";

@Injectable()
export class UserService {
  constructor(private _reqService: RequestService) { }
  _api_url = 'user/';

  list(query: any, page: number, pageSize: number) {
    _.merge(query, {page: page, pageSize: pageSize});
    return this._reqService.post(this._api_url + 'list', query);
  }

  get(id: number) {
    let model = {id: id};
    return this._reqService.post(this._api_url + 'get', model);
  }

  save(user: any, groups: any[]) {
    return this._reqService.post(this._api_url + 'save', {user: user, relations: groups});
  }

  delete(id: number) {
    let model = {id: id};
    return this._reqService.post(this._api_url + 'delete', model);
  }

}

