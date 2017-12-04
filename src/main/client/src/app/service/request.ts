import * as _ from 'lodash';

import {Injectable} from '@angular/core';
import {Http, Headers, RequestOptions, Response} from '@angular/http';

import {Observable} from 'rxjs/Observable';
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/catch';

import {CONSTANT} from '../utils/constant';
import { RouteService } from './route';

@Injectable()
export class RequestService {

    constructor(private http: Http, private routeService: RouteService) {

    }
    post(apiPath: string, reqBody: any) {
        let me = this;

        let url = CONSTANT.API_URL + apiPath;

        let body = JSON.stringify(reqBody);
        let headers = new Headers({ 'Content-Type': 'application/json',
                'token': CONSTANT.TOKEN });
        let options = new RequestOptions({ headers: headers, withCredentials: true });

        console.log(url, body);
        return this.http.post(url, body, options)
            .map(
                function(res) {
                    let json = res.json();
                    console.log(json);
                    if (!!json.code && json.code > 0) {

                    } else if (json.code == -100) {
                      me.routeService.navTo('/login');
                    } else {
                        me.handleError(json.msg);
                    }
                    return json;
                }
            )
            .catch(this.handleError);
    }

    get(apiPath: string) {
      let me = this;
      let url = CONSTANT.API_URL + apiPath;

      console.log(url);
      let headers = new Headers({ 'Content-Type': 'application/json' });
      let options = new RequestOptions({ headers: headers });

      return this.http.get(url, options)
        .map(
          function(res) {
            let json = res.json();
            console.log(json);
            if (!!json.code && json.code > 0) {

            } else if (json.code == -100) {
              me.routeService.navTo('/login');
            } else {
              me.handleError(json.msg);
            }
            return json;
          }
        )
        .catch(this.handleError);
    }

    delete(apiPath: string) {
      let me = this;
      let url = CONSTANT.API_URL + apiPath;

      console.log(url);
      let headers = new Headers({ 'Content-Type': 'application/json', 'token': 'test' });
      let options = new RequestOptions({ headers: headers });

      return this.http.delete(url, options)
        .map(
          function(res) {
            let json = res.json();
            console.log(json);

            if (!!json.code && json.code > 0) {

            } else if (json.code == -100) {
              me.routeService.navTo('/login');
            } else {
              me.handleError(json.msg);
            }
            return json;
          }
        )
        .catch(this.handleError);
    }

    handleError(error: string) {
        console.error(error);
        return Observable.throw(error || 'Server error');
    }
}
