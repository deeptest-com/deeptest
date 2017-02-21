import { Inject, Injectable } from '@angular/core';
import { Observable } from 'rxjs/Observable';
import { Http, Headers } from '@angular/http';
import 'rxjs/add/operator/map';

import {PostService} from './post';

@Injectable()
export class ExchangeService {
    static ENDPOINT: string = '/exchange/index';

    constructor(private _postService: PostService) {
    }

    getData(req): Observable<any> {
        return this._postService.post(ExchangeService.ENDPOINT, req);
    }
}