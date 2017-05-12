import * as _ from 'lodash';

import {Injectable} from '@angular/core';

import {CONSTANT} from '../utils/constant';
import {RequestService} from './request';

@Injectable()
export class OrganizerService {
    constructor(private _reqService: RequestService) { }
    _api_url = 'organizer/';

    list(eventId: number) {
        return this._reqService.post(this._api_url + 'list', {eventId: eventId});
    }

    get(id: number) {
        return this._reqService.post(this._api_url + 'get', {eventId: id});
    }

    save(model: number) {
        return this._reqService.post(this._api_url + 'save', model);
    }

    remove(id: number) {
      return this._reqService.post(this._api_url + 'remove', {id: id});
    }
}

