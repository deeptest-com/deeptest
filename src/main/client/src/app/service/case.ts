import * as _ from "lodash";

import {Injectable} from "@angular/core";

import {CONSTANT} from "../utils/constant";
import {RequestService} from "./request";

@Injectable()
export class CaseService {
  constructor(private _reqService: RequestService) {
  }

  _api_url = 'case/';

  query(projectId: number) {
    return this._reqService.post(this._api_url + 'query', {projectId: projectId});
  }

  queryForSelection(projectId: number, runId: number) {
    return this._reqService.post(this._api_url + 'queryForSelection', {projectId: projectId, runId: runId});
  }

  get(id: number) {
    let model = {id: id};
    return this._reqService.post(this._api_url + 'get', model);
  }

  rename(projectId: number, model: any) {
    _.merge(model, {projectId: projectId});
    return this._reqService.post(this._api_url + 'rename', model);
  }
  move(projectId: number, data: any) {
    _.merge(data, {projectId: projectId});

    return this._reqService.post(this._api_url + 'move', data);
  }
  delete(id: any) {
    let model = {id: id};
    return this._reqService.post(this._api_url + 'delete', model);
  }

  save(projectId: number, model: any) {
    let data = _.clone(model);
    data.steps = null;
    _.merge(data, {projectId: projectId})
    return this._reqService.post(this._api_url + 'save', data);
  }

  saveField(id: number, field: any) {
    let model = _.merge(field, {id: id});
    return this._reqService.post(this._api_url + 'saveField', model);
  }

  changeContentType(contentType: string, id: number) {
    return this._reqService.post(this._api_url + 'changeContentType', {id: id, contentType: contentType});
  }

  reviewPass(id: number, pass: boolean) {
    return this._reqService.post(this._api_url + 'reviewPass', {id: id, pass: pass});
  }
}
