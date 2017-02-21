import {Injectable} from '@angular/core';
import {Http, Headers, RequestOptions, Response} from '@angular/http';

import {Observable} from 'rxjs/Observable';
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/catch';

import {CONSTANT} from '../utils/constant';

@Injectable()
export class PostService {

    constructor(private http: Http) {

    }
    post(apiPath: string, reqBody: any) {
        let me = this;
        let url = CONSTANT.SERVICE_URL + 'api/client/' + CONSTANT.API_VER + apiPath;

        console.log(" reqUrl: ", url);
        console.log("reqBody: ",  reqBody);

        let body = JSON.stringify(reqBody);
        let headers = new Headers({ 'Content-Type': 'application/json', 'token': CONSTANT.TOKEN });
        let options = new RequestOptions({ headers: headers });

        return this.http.post(url, body, options)
            .map(
                function(res) {
                    let json = res.json();
                    console.log("resBody: ",  json);
                    
                    if (!!json.code && json.code > 0) {
                        return json;
                    } else {
                        me.handleError(json.msg);
                    }
                }
            )
            .catch(this.handleError);
    }

    get(apiPath: string) {
      let me = this;
      let url = CONSTANT.SERVICE_URL + 'api/' + CONSTANT.API_VER + apiPath;

      console.log(url);
      let headers = new Headers({ 'Content-Type': 'application/json', 'token': CONSTANT.TOKEN });
      let options = new RequestOptions({ headers: headers });

      return this.http.get(url, options)
        .map(
          function(res) {
            let json = res.json();
            console.log(json);
            if (!!json.code && json.code > 0) {
              return json;
            } else {
              me.handleError(json.msg);
            }
          }
        )
        .catch(this.handleError);
    }

    delete(apiPath: string) {
      let me = this;
      let url = CONSTANT.SERVICE_URL + 'api/' + CONSTANT.API_VER + apiPath;

      console.log(url);
      let headers = new Headers({ 'Content-Type': 'application/json', 'token': 'test' });
      let options = new RequestOptions({ headers: headers });

      return this.http.delete(url, options)
        .map(
          function(res) {
            let json = res.json();
            console.log(json);
            if (!!json.code && json.code > 0) {
              return json;
            } else {
              me.handleError(json.msg);
            }
          }
        )
        .catch(this.handleError);
    }

    handleError(error: Response) {
        console.error(error);
        return Observable.throw(error.json().error || 'Server error');
    }
}
