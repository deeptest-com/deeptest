import * as _ from "lodash";

import {Injectable} from "@angular/core";

import {CONSTANT} from "../utils/constant";
import {RequestService} from "./request";

@Injectable()
export class MsgService {
  constructor(private _reqService: RequestService) {
  }

  _api_url = 'msg/';

  list(query: any, page: number, pageSize: number) {
    _.merge(query, {page: page, pageSize: pageSize});
    return this._reqService.post(this._api_url + 'list', query);
  }

  get(id: number) {
    let model = {id: id};
    return this._reqService.post(this._api_url + 'get', model);
  }

  delete(id: any) {
    let model = {id: id};
    return this._reqService.post(this._api_url + 'delete', model);
  }

  markAllRead() {
    return this._reqService.post(this._api_url + 'markAllRead', {});
  }
  markRead(id: number) {
    return this._reqService.post(this._api_url + 'markRead', {id: id});
  }
}


