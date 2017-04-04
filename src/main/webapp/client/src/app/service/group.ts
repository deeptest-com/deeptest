import {Injectable} from "@angular/core";

import {RequestService} from "./request";

@Injectable()
export class GroupService {
  constructor(private _reqService: RequestService) { }
  _api_url = 'group/';

  list(query: any) {
    return this._reqService.post(this._api_url + 'list', query);
  }

  get(id: number) {
    let model = {id: id};
    return this._reqService.post(this._api_url + 'get', model);
  }

  save(group: any, users: any[]) {
    return this._reqService.post(this._api_url + 'save', {group: group, relations: users});
  }

  delete(id: number) {
    let model = {id: id};
    return this._reqService.post(this._api_url + 'delete', model);
  }

  listByUser(id: number) {
    return this._reqService.post(this._api_url + 'listByUser', {userId: id});
  }
  saveByUser(userId: number, models: any) {
    return this._reqService.post(this._api_url + 'saveByUser', {userId: userId, models: models});
  }
}

