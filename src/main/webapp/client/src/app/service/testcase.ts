import {Injectable} from '@angular/core';

import {CONSTANT} from '../utils/constant';
import {RequestService} from './request';

@Injectable()
export class TestcaseService {
    constructor(private _reqService: RequestService) { }
    _api_url = 'testcase/';

    query(query: any) {
        return this._reqService.post(this._api_url + 'query', query);
    }

    get(id: number) {
        return this._reqService.post(this._api_url + 'get', {eventId: id});
    }

    save(model: number) {
        return this._reqService.post(this._api_url + 'save', model);
    }
}

