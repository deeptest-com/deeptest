import * as _ from 'lodash';

import {Injectable} from "@angular/core";
import {GlobalState} from "../global.state";
import {RequestService} from "./request";

@Injectable()
export class ProjectService {
  constructor(private _reqService:RequestService, private _state:GlobalState) {
  }

  _api_url = 'project/';

  list(query:any) {
    return this._reqService.post(this._api_url + 'list', query);
  }

  get(id:number) {
    let model = {id: id};
    return this._reqService.post(this._api_url + 'get', model);
  }

  save(model:any) {
    return this._reqService.post(this._api_url + 'save', {model: model});
  }
  saveMembers(model:any, entityTypeAndIds: string[]) {
    _.merge(model, {entityTypeAndIds: entityTypeAndIds});
    return this._reqService.post(this._api_url + 'saveMembers', model);
  }
  changeRole(projectId: number, roleId: number, entityId: number) {
    return this._reqService.post(this._api_url + 'changeRole', {projectId: projectId, roleId: roleId, entityId: entityId});
  }

  delete(id:number) {
    let model = {id: id};
    return this._reqService.post(this._api_url + 'delete', model);
  }

  view(projectId: number) {
    let model = {id: projectId};
    return this._reqService.post(this._api_url + 'view', model);
  }
}

