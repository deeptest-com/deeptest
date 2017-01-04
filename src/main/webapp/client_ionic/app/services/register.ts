import { Inject, Injectable } from '@angular/core';
import { Observable } from 'rxjs/Observable';
import { Http, Headers } from '@angular/http';
import 'rxjs/add/operator/map';

import {PostService} from './post';

@Injectable()
export class RegisterService {
    static ENDPOINT_INDEX: string = '/register/getInfo';
    static ENDPOINT_REGISTER: string = '/register/register';

    constructor(private _postService: PostService) {
    }

    getInfo(req): Observable<any> {
      return this._postService.post(RegisterService.ENDPOINT_INDEX, req);
    }

    register(req): Observable<any> {
        return this._postService.post(RegisterService.ENDPOINT_REGISTER, req);
    }
}
