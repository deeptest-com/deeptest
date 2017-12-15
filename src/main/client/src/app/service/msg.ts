import * as _ from "lodash";

import {Injectable} from "@angular/core";

import {CONSTANT} from "../utils/constant";
import {RequestService} from "./request";

@Injectable()
export class MsgService {
  constructor(private _reqService: RequestService) {
  }

  _api_url = 'msg/';

  list(userId: number) {
    return this._reqService.post(this._api_url + 'list', {userId: userId});
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

}


