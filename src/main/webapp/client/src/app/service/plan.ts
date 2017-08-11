import * as _ from "lodash";

import {Injectable} from "@angular/core";

import {CONSTANT} from "../utils/constant";
import {RequestService} from "./request";

import {TreeModel} from "../components/ng2-tree";

@Injectable()
export class PlanService {
  constructor(private _reqService: RequestService) {
  }

  _api_url = 'plan/';

  query(projectId: number, query: any) {
    _.merge(query, {projectId: projectId});
    return this._reqService.post(this._api_url + 'query', query);
  }

  get(id: number) {
    let model = {id: id};
    return this._reqService.post(this._api_url + 'get', model);
  }

  save(model: number) {
    return this._reqService.post(this._api_url + 'save', model);
  }

  delete(id: number) {
    let model = {id: id};
    return this._reqService.post(this._api_url + 'delete', model);
  }
}


