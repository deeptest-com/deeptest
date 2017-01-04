import { Inject, Injectable } from '@angular/core';
import { Observable } from 'rxjs/Observable';
import { Http, Headers } from '@angular/http';
import 'rxjs/add/operator/map';

import {PostService} from './post';

@Injectable()
export class EventService {
    static GET_EVENT: string = '/event/getEvent';
    static GET_DETAIL: string = '/event/getDetail';

    constructor(private _postService: PostService) {
    }

    getEvent(req): Observable<any> {
        return this._postService.post(EventService.GET_EVENT, req);
    }
    getDetail(req): Observable<any> {
        return this._postService.post(EventService.GET_DETAIL, req);
    }
}
