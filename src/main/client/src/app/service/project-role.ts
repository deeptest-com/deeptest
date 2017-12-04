import * as _ from 'lodash';

import {Injectable} from "@angular/core";

import { Cookie } from 'ng2-cookies/ng2-cookies';
import {GlobalState} from '../global.state';

import { CONSTANT } from '../utils/constant';
import { RouteService } from './route';
import {RequestService} from "./request";

@Injectable()
export class ProjectRoleService {
  constructor(private _reqService: RequestService) { }
  _api_url = 'project_role/';

  list(query: any) {
    return this._reqService.post(this._api_url + 'list', query);
  }

  get(id: number) {
    let model = {id: id};
    return this._reqService.post(this._api_url + 'get', model);
  }

  save(projectRole: any, projectPrivileges: any[]) {
    return this._reqService.post(this._api_url + 'save', {projectRole: projectRole, projectPrivileges: projectPrivileges});
  }

  delete(id: number) {
    let model = {id: id};
    return this._reqService.post(this._api_url + 'delete', model);
  }
}

