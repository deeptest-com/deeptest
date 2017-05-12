import * as _ from 'lodash';

import {Injectable} from '@angular/core';

import {CONSTANT} from '../utils/constant';
import {RequestService} from './request';

@Injectable()
export class ScheduleService {
    constructor(private _reqService: RequestService) { }
    _api_url = 'schedule/';

    listByEvent(eventId: number) {
        return this._reqService.post(this._api_url + 'list', {eventId: eventId, isNest: true});
    }

    get(id: number) {
        return this._reqService.post(this._api_url + 'get', {eventId: id});
    }

    save(model: any) {
        return this._reqService.post(this._api_url + 'save', model);
    }

}

