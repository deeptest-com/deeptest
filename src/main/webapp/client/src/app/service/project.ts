import {Injectable} from '@angular/core';

import {CONSTANT} from '../utils/constant';
import {RequestService} from './request';

@Injectable()
export class ProjectService {
    constructor(private _reqService: RequestService) { }
    _api_url = 'project/';

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
}

