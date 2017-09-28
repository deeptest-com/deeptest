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
  getUsers(projectId: number) {
    let model = {projectId: projectId};
    return this._reqService.post(this._api_url + 'getUsers', model);
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

  setSize(left: number, right: number) {
    let model = {left: left, right: right};
    return this._reqService.post(this._api_url + 'setSize', model);
  }

  search(orgId:number, keywords: string, exceptIds: string[]) {
    let model = {orgId:orgId, keywords: keywords, exceptIds: exceptIds};
    return this._reqService.post(this._api_url + 'search', model);
  }

}

