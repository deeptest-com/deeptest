import * as _ from "lodash";

import {Injectable} from "@angular/core";

import {CONSTANT} from "../utils/constant";
import {RequestService} from "./request";

@Injectable()
export class RunService {
  constructor(private _reqService: RequestService) {
  }

  _api_url = 'run/';

  loadCase(projectId: number, runId: number) {
    return this._reqService.post(this._api_url + 'loadCase', {projectId: projectId, runId: runId});
  }

  get(id: number) {
    let model = {id: id};
    return this._reqService.post(this._api_url + 'get', model);
  }

  saveRun(prjId: number, planId: number, run: any) {
    return this._reqService.post(this._api_url + 'save',
      {prjId: prjId, planId: planId, id: run.id, name: run.name, userId: run.userId});
  }

  saveRunCases(planId: number, runId: number, cases: any[]) {
    let ids: number[] = cases.map(function (item,index,input) {
      return item.id;
    });
    return this._reqService.post(this._api_url + 'saveCases', {planId: planId, runId: runId, cases: ids});
  }

  delete(id: any) {
    let model = {id: id};
    return this._reqService.post(this._api_url + 'delete', model);
  }
  close(id: any) {
    let model = {id: id};
    return this._reqService.post(this._api_url + 'close', model);
  }

  markAllRead(ids: number[]) {
    return this._reqService.post(this._api_url + 'markAllRead', {ids: ids});
  }

}


