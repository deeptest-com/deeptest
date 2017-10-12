import * as _ from "lodash";

import {Injectable} from "@angular/core";

import {CONSTANT} from "../utils/constant";
import {RequestService} from "./request";

import {TreeModel} from "../components/ng2-tree";

@Injectable()
export class CaseInRunService {
  constructor(private _reqService: RequestService) {
  }

  _api_url = 'caseInRun/';

  query(runId: number) {
    return this._reqService.post(this._api_url + 'query', {runId: runId});
  }

  get(id: number) {
    let model = {id: id};
    return this._reqService.post(this._api_url + 'get', model);
  }
}


