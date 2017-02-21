import { Inject, Injectable } from '@angular/core';
import { Observable } from 'rxjs/Observable';
import { Http, Headers } from '@angular/http';
import 'rxjs/add/operator/map';

import {PostService} from './post';

@Injectable()
export class GuestService {
    static ENDPOINT_LIST: string = '/guest/list';

    constructor(private _postService: PostService) {
    }

    list(req): Observable<any> {
        return this._postService.post(GuestService.ENDPOINT_LIST, req);
    }
}