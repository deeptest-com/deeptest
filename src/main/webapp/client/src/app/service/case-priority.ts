import {Injectable} from "@angular/core";

import {RequestService} from "./request";

@Injectable()
export class CasePriorityService {
  constructor(private _reqService: RequestService) { }
  _api_url = 'case_priority/';

  list(query: any) {
    return this._reqService.post(this._api_url + 'list', query);
  }

  get(id: number) {
    let model = {id: id};
    return this._reqService.post(this._api_url + 'get', model);
  }

  save(model: any) {
    return this._reqService.post(this._api_url + 'save', {model: model});
  }

  delete(id: number) {
    let model = {id: id};
    return this._reqService.post(this._api_url + 'delete', model);
  }

  setDefault(id: number) {
    let model = {id: id};
    return this._reqService.post(this._api_url + 'setDefault', model);
  }

}

