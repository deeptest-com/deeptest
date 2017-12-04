import * as _ from 'lodash';

import {Injectable} from '@angular/core';

import {CONSTANT} from '../utils/constant';
import {RequestService} from './request';

@Injectable()
export class CompanyService {
    constructor(private _reqService: RequestService) { }
    _api_url = 'company/';

    list(pageSize: number, page: number, eventId: number) {
        return this._reqService.post(this._api_url + 'list', {pageSize: pageSize, page: page, eventId: eventId});
    }

    get() {
        return this._reqService.post(this._api_url + 'get', {});
    }

    save(model: number) {
        return this._reqService.post(this._api_url + 'save', model);
    }

    remove(id: number) {
      return this._reqService.post(this._api_url + 'remove', {id: id});
    }
}

