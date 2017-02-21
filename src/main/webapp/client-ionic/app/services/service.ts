import { Inject, Injectable } from '@angular/core';
import { Observable } from 'rxjs/Observable';
import { Http, Headers } from '@angular/http';
import 'rxjs/add/operator/map';

import {PostService} from './post';

@Injectable()
export class ServiceService {
    static ENDPOINT_LIST: string = '/service/list';

    constructor(private _postService: PostService) {
    }

    getData(req): Observable<any> {
        return this._postService.post(ServiceService.ENDPOINT_LIST, req);
    }
}
