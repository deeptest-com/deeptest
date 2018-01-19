import * as _ from "lodash";

import {Injectable} from "@angular/core";

import {CONSTANT} from "../utils/constant";
import {RequestService} from "./request";

@Injectable()
export class AlertService {
  constructor(private _reqService: RequestService) {
  }

  _api_url = 'alert/';

  markAllRead(ids: string) {
    return this._reqService.post(this._api_url + 'markAllRead', {ids: ids});
  }

}


