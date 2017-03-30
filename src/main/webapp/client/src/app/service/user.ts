import {Injectable} from "@angular/core";

import {RequestService} from "./request";

@Injectable()
export class UserService {
  constructor(private _reqService: RequestService) { }
  _api_url = 'user/';

  list(query: any) {
    return this._reqService.post(this._api_url + 'list', query);
  }

  get(id: number) {
    let model = {id: id};
    return this._reqService.post(this._api_url + 'get', model);
  }

  save(model: any) {
    return this._reqService.post(this._api_url + 'save', model);
  }

  delete(id: number) {
    let model = {id: id};
    return this._reqService.post(this._api_url + 'delete', model);
  }

  saveGroups(model: any) {
    return this._reqService.post(this._api_url + 'saveGroups', {model: model});
  }

}

