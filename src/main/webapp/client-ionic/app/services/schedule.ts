import { Inject, Injectable } from '@angular/core';
import { Observable } from 'rxjs/Observable';
import { Http, Headers } from '@angular/http';
import 'rxjs/add/operator/map';

import {PostService} from './post';

@Injectable()
export class ScheduleService {
    static ENDPOINT_LIST: string = '/schedule/listByEvent';
    static ENDPOINT_SESSIONS: string = '/schedule/listSession';

    constructor(private _postService: PostService) {
    }

    getData(req): Observable<any> {
        return this._postService.post(ScheduleService.ENDPOINT_LIST, req);
    }
    
    getSessions(req): Observable<any> {
        return this._postService.post(ScheduleService.ENDPOINT_SESSIONS, req);
    }
}