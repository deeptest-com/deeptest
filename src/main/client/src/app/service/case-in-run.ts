import * as _ from "lodash";

import {Injectable} from "@angular/core";

import {CONSTANT} from "../utils/constant";
import {RequestService} from "./request";

@Injectable()
export class CaseInRunService {
  constructor(private _reqService: RequestService) {
  }

  _api_url = 'caseInRun/';

  query(projectId:number , runId: number) {
    return this._reqService.post(this._api_url + 'query', {projectId: projectId, runId: runId});
  }

  get(id: number) {
    let model = {id: id};
    return this._reqService.post(this._api_url + 'get', model);
  }

  setResult(modelId: any, result: string, nextId: number, status: string) {
    let data = {id: modelId, result: result, nextId: nextId, status: status};
    return this._reqService.post(this._api_url + 'setResult', data);
  }

}
