import * as _ from 'lodash';

import {Injectable} from "@angular/core";

import {Observable} from 'rxjs/Observable';
import 'rxjs/add/observable/of';

import { Cookie } from 'ng2-cookies/ng2-cookies';
import {GlobalState} from '../global.state';

import { CONSTANT } from '../utils/constant';
import { RouteService } from './route';
import {RequestService} from "./request";

@Injectable()
export class AccountService {
  constructor(private _state:GlobalState, private _reqService:RequestService, private routeService: RouteService) {
  }

  _login = 'account/login';
  _logout = 'account/logout';
  _register = 'account/register';
  _changePassword = 'account/changePassword';

  _forgotPassword = 'account/forgotPassword';
  _checkResetPassword = 'account/checkResetPassword';
  _resetPassword = 'account/resetPassword';

  _getProfile = 'account/getProfile';
  _getInfo = 'account/getInfo';
  _saveProfile = 'account/saveProfile';
  _saveInfo = 'account/saveInfo';
  _suggestions = 'suggestions/:id';

  _collections = 'collections/:id';
  _removeCollection = 'user/removeCollections';
  _msgs = 'msgs';

  login(model: any) {
    let that = this;
    return this._reqService.post(this._login, model).map((json:any) => {
      let errors = undefined;
      if (json.code == 1) {
        let days:number = model.rememberMe? 30: 1;

        that.saveTokenLocal(json.token, days);
        that.routeService.navTo('/pages/org/' + json.profile.defaultOrgId + '/prjs');
      } else {
        errors = json.msg;
      }
      return errors;
    });
  }
  register(model: any) {
    let that = this;
    return this._reqService.post(this._register, model).map((json:any) => {
      let errors = undefined;
      if (json.code == 1) {
        that.saveTokenLocal(json.token, 1);

        // that.changeProfile(json.profile);
        // that.changeRecentProject(json.recentProjects);

        that.routeService.navTo('/pages/dashboard');
      } else {
        errors = json.msg;
      }
      return errors;
    });
  }

  resetPassword(vcode: string, model:number) {
    _.merge(model, {vcode: vcode});
    return this._reqService.post(this._resetPassword, model).map((json:any) => {
      let errors = undefined;
      if (json.code == 1) {
        this.saveTokenLocal(json.token, 1);

        // that.changeProfile(json.profile);
        // that.changeRecentProject(json.recentProjects);

        this.routeService.navTo('/pages/dashboard');
      } else {
        errors = json.data;
      }
      return errors;
    });
  }

  checkResetPassword(vcode: string) {
    return this._reqService.post(this._checkResetPassword, {vcode: vcode});
  }

  loadProfileRemote(): Observable<any> {
    let that = this;
    let token = Cookie.get(CONSTANT.TOKEN_KEY);
    console.log('token from cookie: ', token);

    if (token) {
      CONSTANT.TOKEN = JSON.parse(token);

      return this._reqService.post(that._getProfile, {}).map(json => {
        if (json.code == 1) {
          that.changeProfile(json.profile);
          that.changeMyOrgs(json.profile.orgs, json.profile.defaultOrgId);
          that.changeRecentProjects(json.profile.recentProjects);
          that.changeCasePropertyMap(json.profile.casePropertyMap);

          return Observable.of(true);
        } else {
          return Observable.of(false);
        }
      });
    } else  {
      return Observable.of(false);
    }
  }

  logout() {
    this._reqService.post(this._logout, {}).subscribe((json:any) => {
      // if (json.code == 1) {
        Cookie.delete(CONSTANT.TOKEN_KEY);
        // CONSTANT.PROFILE = null;
        this.routeService.navTo('/login');
      // }
    });
  }

  forgotPassword(email:number) {
    return this._reqService.post(this._forgotPassword, {email: email});
  }

  getProfile() {
    return this._reqService.post(this._getProfile, {});
  }
  getInfo() {
    return this._reqService.post(this._getInfo, {});
  }

  saveProfile(profile:any) {
    return this._reqService.post(this._saveProfile, profile);
  }
  saveInfo(profile:any) {
    return this._reqService.post(this._saveInfo, profile);
  }

  changePassword(model:any) {
    return this._reqService.post(this._changePassword, model);
  }

  saveSuggestion(content) {
    return this._reqService.post(this._suggestions.replace(':id', ''), {suggestion: {content: content}});
  }

  saveTokenLocal(token: any, expireDays: number) {
    let that = this;
    CONSTANT.TOKEN = token;

    if (!expireDays) {
      expireDays = parseInt(Cookie.get(CONSTANT.TOKEN_EXPIRE));
    } else {
      Cookie.set(CONSTANT.TOKEN_EXPIRE, expireDays + '', 365);
    }

    Cookie.set(CONSTANT.TOKEN_KEY, JSON.stringify(token), expireDays);
  }

  changeProfile(profile: any) {
    CONSTANT.PROFILE = profile;
    this._state.notifyDataChanged(CONSTANT.STATE_CHANGE_PROFILE, profile);
  }

  changeMyOrgs(orgs: any[], currOrgId: number, gotoDefault: boolean = false) {
    if (orgs) {
      CONSTANT.ALL_ORGS = orgs;
    }
    if (currOrgId) {
      CONSTANT.CURR_ORG_ID = currOrgId;
      console.log('change orgId ' + CONSTANT.CURR_ORG_ID);
    }
    if (gotoDefault) {this.routeService.navTo("/pages/project/list");}

    this._state.notifyDataChanged(CONSTANT.STATE_CHANGE_ORGS, {orgs: orgs, currOrgId: currOrgId});
  }

  changeRecentProjects(recentProjects: any[]) {
    CONSTANT.RECENT_PROJECTS = recentProjects;

    if (recentProjects.length > 0) {
      CONSTANT.CURR_PRJ_ID = recentProjects[0].projectId;
      CONSTANT.CURR_PRJ_NAME = recentProjects[0].projectName;
    } else {
      CONSTANT.CURR_PRJ_ID = undefined;
      CONSTANT.CURR_PRJ_NAME = undefined;
    }

    this._state.notifyDataChanged(CONSTANT.STATE_CHANGE_PROJECTS,
      {recentProjects: CONSTANT.RECENT_PROJECTS});
  }

  changeCasePropertyMap(casePropertyMap: any) {
    CONSTANT.CASE_PROPERTY_MAP = casePropertyMap;
  }

}
