import { Inject, Injectable } from '@angular/core';
import { Observable } from 'rxjs/Observable';
import { Http, Headers } from '@angular/http';
import 'rxjs/add/operator/map';

import {PostService} from './post';

@Injectable()
export class FeedbackService {
    static ENDPOINT_LIST: string = '/feadback/list';

    constructor(private _postService: PostService) {
    }

    list(req): Observable<any> {
        return this._postService.post(FeedbackService.ENDPOINT_LIST, req);
    }
}
