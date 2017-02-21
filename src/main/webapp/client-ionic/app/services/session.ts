import { Inject, Injectable } from '@angular/core';
import { Observable } from 'rxjs/Observable';
import { Http, Headers } from '@angular/http';
import 'rxjs/add/operator/map';

import {PostService} from './post';

@Injectable()
export class SessionService {
    static ENDPOINT: string = '/session/';

    constructor(private _postService: PostService) {
      
    }

}
