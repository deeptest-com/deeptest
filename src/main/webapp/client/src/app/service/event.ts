import {Injectable} from '@angular/core';

import {CONSTANT} from '../utils/constant';
import {RequestService} from './request';

@Injectable()
export class EventService {
    constructor(private _reqService: RequestService) { }
    _api_url = 'event/';

    list(itemsPerPage: number, currentPage: number, status: string) {
        return this._reqService.post(this._api_url + 'list', {itemsPerPage: itemsPerPage, currentPage: currentPage, status: status});
    }

    get(id: number) {
        return this._reqService.post(this._api_url + 'get', {eventId: id});
    }

    save(model: number) {
        return this._reqService.post(this._api_url + 'save', model);
    }
}

