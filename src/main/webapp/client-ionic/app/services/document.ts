import { Inject, Injectable } from '@angular/core';
import { Observable } from 'rxjs/Observable';
import { Http, Headers } from '@angular/http';
import 'rxjs/add/operator/map';

import {PostService} from './post';

@Injectable()
export class DocumentService {
    static ENDPOINT: string = '/document/list';

    constructor(private _postService: PostService) {
    }

    list(req): Observable<any> {
        return this._postService.post(DocumentService.ENDPOINT, req);
    }
}