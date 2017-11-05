import * as _ from "lodash";

import {Injectable} from "@angular/core";

import {CONSTANT} from "../utils/constant";
import {RequestService} from "./request";

@Injectable()
export class SuiteService {
  constructor(private _reqService: RequestService) {
  }

  _api_url = 'suite/';

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

  create(node: any) {
    let model = {id: node.id, value: node.value, type: node.type, pid: node.pid};
    return this._reqService.post(this._api_url + 'create', model);
  }

  move(target: any, src: any, options: any) {
    let model;
    if (options.mode === 'inner') {
      model = {id: src.id, newPid: target.id, prePid: src.pid};
    } else {
      model = {id: src.id, newPid: target.pid, prePid: src.pid};
    }
    _.merge(model, options);
    return this._reqService.post(this._api_url + 'move', model);
  }

  rename(node: any) {
    console.log('rename');
    let model = {id: node.id, value: node.value, type: node.type, pid: node.pid};
    return this._reqService.post(this._api_url + 'rename', model);
  }

  delete(node: any) {
    let model = {id: node.id};
    return this._reqService.post(this._api_url + 'delete', model);
  }
}


