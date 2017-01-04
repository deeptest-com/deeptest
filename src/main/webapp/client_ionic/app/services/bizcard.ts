import { Inject, Injectable } from '@angular/core';
import { Observable } from 'rxjs/Observable';
import { Http, Headers } from '@angular/http';
import 'rxjs/add/operator/map';

import {PostService} from './post';

@Injectable()
export class BizcardService {
    static ENDPOINT_GET: string = '/bizcard/getBizcard';
    static ENDPOINT_SAVE: string = '/bizcard/save';

    constructor(private _postService: PostService) {
    }

    getMyBizcard(req): Observable<any> {
        return this._postService.post(BizcardService.ENDPOINT_GET, req);
    }

    save(req): Observable<any> {
      return this._postService.post(BizcardService.ENDPOINT_SAVE, req);
    }
}
