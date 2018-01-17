import * as _ from 'lodash';

import {Injectable} from "@angular/core";
import {Observable} from 'rxjs/Observable';

import {RequestService} from "./request";

import { Cookie } from 'ng2-cookies/ng2-cookies';
import { CONSTANT } from '../utils/constant';
import {GlobalState} from '../global.state';
import { RouteService } from './route';

@Injectable()
export class UserService {
  constructor(private _state:GlobalState, private _routeService: RouteService, private _reqService: RequestService) { }
  _api_url = 'user/';

  _getProfile = this._api_url + 'getProfile';
  _saveInfo = this._api_url + 'saveInfo';
  _setLeftSize = this._api_url + 'setLeftSize';

  list(query: any, page: number, pageSize: number) {
    _.merge(query, {page: page, pageSize: pageSize});
    return this._reqService.post(this._api_url + 'list', query);
  }
  getUsers(projectId: number) {
    let model = {projectId: projectId};
    return this._reqService.post(this._api_url + 'getUsers', model);
  }

  get(id: number) {
    let model = {id: id};
    return this._reqService.post(this._api_url + 'get', model);
  }

  save(user: any, groups: any[]) {
    return this._reqService.post(this._api_url + 'save', {user: user, relations: groups});
  }
  invite(user: any, groups: any[]) {
    return this._reqService.post(this._api_url + 'invite', {user: user, relations: groups});
  }

  delete(id: number) {
    let model = {id: id};
    return this._reqService.post(this._api_url + 'delete', model);
  }

  search(orgId:number, keywords: string) {
    let model = {orgId:orgId, keywords: keywords};
    return this._reqService.post(this._api_url + 'search', model);
  }

  loadProfileRemote(context = {}): Observable<any> {
    let that = this;
    let token = Cookie.get(CONSTANT.TOKEN_KEY);
    console.log('token from cookie: ', token);

    if (token && token != 'undefined') {
      CONSTANT.TOKEN = JSON.parse(token);

      return this._reqService.post(that._getProfile, context).map(json => {
        if (json.code == 1) {
          CONSTANT.CURR_ORG_ID = json.profile.defaultOrgId;
          CONSTANT.CURR_PRJ_ID = json.profile.defaultPrjId;
          CONSTANT.CURR_PRJ_NAME = json.profile.defaultPrjName;

          CONSTANT.PROFILE = json.profile;
          CONSTANT.SYS_PRIVILEGES = json.sysPrivileges;
          CONSTANT.MY_ORGS = json.myOrgs;
          CONSTANT.ORG_PRIVILEGES = json.orgPrivileges;
          CONSTANT.CASE_PROPERTY_MAP = json.casePropertyMap;

          CONSTANT.RECENT_PROJECTS = json.recentProjects;
          CONSTANT.PRJ_PRIVILEGES = json.prjPrivileges;

          return Observable.of(true);
        } else {
          this._routeService.navTo('/login');
          return Observable.of(false);
        }
      });
    } else  {
      this._routeService.navTo('/login');
      return Observable.of(false);
    }
  }

  saveInfo(profile:any) {
    return this._reqService.post(this._saveInfo, profile);
  }

  setLeftSize(left: any) {
    let model = {left: left};
    return this._reqService.post(this._setLeftSize, model);
  }



}

