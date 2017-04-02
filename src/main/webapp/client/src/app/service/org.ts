import {Injectable} from "@angular/core";

import {RequestService} from "./request";

@Injectable()
export class OrgService {
  constructor(private _reqService: RequestService) { }
  _api_url = 'org/';

  list(query: any) {
    return this._reqService.post(this._api_url + 'list', query);
  }

  get(id: number) {
    let model = {id: id};
    return this._reqService.post(this._api_url + 'get', model);
  }

  save(user: any, groups: any[]) {
    return this._reqService.post(this._api_url + 'save', {user: user, groups: groups});
  }

  delete(id: number) {
    let model = {id: id};
    return this._reqService.post(this._api_url + 'delete', model);
  }

}

