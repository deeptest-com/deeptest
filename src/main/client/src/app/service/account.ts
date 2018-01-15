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
  constructor(private _state:GlobalState, private _reqService:RequestService, private _routeService: RouteService) {
  }

  _login = 'account/login';
  _loginWithVcode = 'account/loginWithVcode';
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
  _setLeftSize = 'account/setLeftSize';
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
        that._routeService.navTo('/pages/org/' + json.profile.defaultOrgId + '/prjs');
      } else {
        errors = json.msg;
      }
      return errors;
    });
  }
  loginWithVcode(vcode: string) {
    let that = this;
    return this._reqService.post(this._loginWithVcode, {vcode: vcode}).map((json:any) => {
      let errors = undefined;
      if (json.code == 1) {
        let days:number = 1;

        that.saveTokenLocal(json.token, days);
        that._routeService.navTo('/pages/org/' + json.profile.defaultOrgId + '/prjs');
      } else {
        errors = json.msg;
      }
      return errors;
    });
  }
  register(model: any) {
    let that = this;
    return this._reqService.post(this._register, model).map((json:any) => {
      if (json.code == 1) {
        that.saveTokenLocal(json.token, 1);
      }
      return json;
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

        this._routeService.navTo('/pages/dashboard');
      } else {
        errors = json.data;
      }
      return errors;
    });
  }

  checkResetPassword(vcode: string) {
    return this._reqService.post(this._checkResetPassword, {vcode: vcode});
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

  logout() {
    this._reqService.post(this._logout, {}).subscribe((json:any) => {
      // if (json.code == 1) {
        Cookie.delete(CONSTANT.TOKEN_KEY);
        // CONSTANT.PROFILE = null;
        this._routeService.navTo('/login');
      // }
    });
  }

  forgotPassword(email:number) {
    return this._reqService.post(this._forgotPassword, {email: email});
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

  setLeftSize(left: any) {
    let model = {left: left};
    return this._reqService.post(this._setLeftSize, model);
  }

}
