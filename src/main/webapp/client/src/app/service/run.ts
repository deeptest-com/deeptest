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

  save(model: number) {
    return this._reqService.post(this._api_url + 'save', model);
  }

  delete(node: TreeModel) {
    let model = {id: node.id};
    return this._reqService.post(this._api_url + 'delete', model);
  }
}


