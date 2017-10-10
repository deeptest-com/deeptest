import * as _ from "lodash";

import {Injectable} from "@angular/core";

import {CONSTANT} from "../utils/constant";
import {RequestService} from "./request";

import {TreeModel} from "../components/ng2-tree";

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

  save(planId: number, runId: number, cases: any[]) {
    let ids: number[] = cases.map(function (item,index,input) {
      return item.id;
    });
    return this._reqService.post(this._api_url + 'save', {planId: planId, runId: runId, cases: ids});
  }

  delete(node: TreeModel) {
    let model = {id: node.id};
    return this._reqService.post(this._api_url + 'delete', model);
  }
}


