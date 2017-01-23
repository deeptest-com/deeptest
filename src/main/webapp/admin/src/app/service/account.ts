import {Injectable} from '@angular/core';

import {CONSTANT} from '../utils/constant';
import {RequestService} from './request';

@Injectable()
export class AccountService {
    constructor(private _reqService: RequestService) { }
    _api_url = 'account/';

    list(itemsPerPage: number, currentPage: number, eventId: number) {
        return this._reqService.post(this._api_url + 'list', {itemsPerPage: itemsPerPage, currentPage: currentPage, eventId: eventId});
    }

    get(id: number) {
        return this._reqService.post(this._api_url + 'get', {eventId: id});
    }

    save(model: any) {
        return this._reqService.post(this._api_url + 'save', model);
    }

    remove(id: number) {
      return this._reqService.post(this._api_url + 'remove', {id: id});
    }
}

