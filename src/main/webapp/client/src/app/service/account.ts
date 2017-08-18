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
  _resetPassword = 'account/resetPassword';

  _getProfile = 'account/getProfile';
  _saveProfile = 'account/saveProfile';
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
        // that.changeProfile(json.profile);
        // that.changeRecentProject(json.recentProjects);

        that.routeService.navTo('/pages/project/list');
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

  resetPassword(model:number) {
    let that = this;
    return this._reqService.post(this._resetPassword, model).map((json:any) => {
      let errors = undefined;
      if (json.code == 1) {
        that.saveTokenLocal(json.token, 1);

        // that.changeProfile(json.profile);
        // that.changeRecentProject(json.recentProjects);

        that.routeService.navTo('/pages/dashboard');
      } else {
        errors = json.data;
      }
      return errors;
    });
  }

  loadProfileRemote(): Observable<any> {
    let that = this;
    let token = Cookie.get(CONSTANT.TOKEN_KEY);
    console.log('token from cookie: ', token);

    if (token) {
      CONSTANT.TOKEN = JSON.parse(token);

      return this._reqService.post(that._getProfile, {}).map(json => {
        that.changeProfile(json.profile);
        that.changeRecentProject(json.recentProjects);
        that.changeMyOrgs(json.myOrgs);

        return json;
      });
    } else  {
      //noinspection TypeScriptUnresolvedFunction
      return Observable.of(false);
    }
  }

  logout() {
    this._reqService.post(this._logout, {}).subscribe((json:any) => {
      if (json.code == 1) {
        Cookie.delete(CONSTANT.TOKEN_KEY);

        this.routeService.navTo('/login');
      }
    });
  }

  forgotPassword(email:number) {
    return this._reqService.post(this._forgotPassword, {email: email});
  }

  getProfile() {
    return this._reqService.post(this._getProfile, {});
  }

  saveProfile(profile:any) {
    return this._reqService.post(this._saveProfile, profile);
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
    this._state.notifyDataChanged('profile.refresh', profile);
  }

  changeRecentProject(recentProjects: any[]) {
    CONSTANT.RECENT_PROJECTS = recentProjects;
    if (recentProjects.length > 0) {
      CONSTANT.CURRENT_PROJECT = {id: recentProjects[0].projectId, name: recentProjects[0].projectName};
    } else {
      CONSTANT.CURRENT_PROJECT = {id: null, name: ''};
    }

    this._state.notifyDataChanged('recent.projects.change', recentProjects);
  }

  changeMyOrgs(orgs: any[]) {
    CONSTANT.MY_ORGS = orgs;
    this._state.notifyDataChanged('my.orgs.change', orgs);
  }

}
