import {Injectable} from "@angular/core";

import { Cookie } from 'ng2-cookies/ng2-cookies';
import {GlobalState} from '../global.state';

import { CONSTANT } from '../utils/constant';
import { RouteService } from './route';
import {RequestService} from "./request";

@Injectable()
export class OrgRoleService {
  constructor(private _reqService: RequestService) { }
  _api_url = 'org_role/';

  list(query: any, currentPage: number, itemsPerPage: number) {
    _.merge(query, {currentPage: currentPage, itemsPerPage: itemsPerPage});
    return this._reqService.post(this._api_url + 'list', query);
  }

  get(id: number) {
    let model = {id: id};
    return this._reqService.post(this._api_url + 'get', model);
  }

  save(orgRole: any, orgPrivileges: any[]) {
    return this._reqService.post(this._api_url + 'save', {orgRole: orgRole, orgPrivileges: orgPrivileges});
  }

  delete(id: number) {
    let model = {id: id};
    return this._reqService.post(this._api_url + 'delete', model);
  }
}

