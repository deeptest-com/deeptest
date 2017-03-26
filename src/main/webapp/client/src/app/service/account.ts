import {Injectable} from "@angular/core";

import { Cookie } from 'ng2-cookies/ng2-cookies';
import {GlobalState} from '../global.state';

import { CONSTANT } from '../utils/constant';
import { RouteService } from './route';
import {RequestService} from "./request";

@Injectable()
export class AccountService {
  constructor(private _state:GlobalState, private _reqService:RequestService, private routeService: RouteService) {
  }

  _login = 'user/login';
  _logout = 'user/logout';
  _register = 'user/register';
  _changePassword = 'user/changePassword';

  _forgotPassword = 'user/forgotPassword';
  _resetPassword = 'user/resetPassword';

  _getProfile = 'user/getProfile';
  _saveProfile = 'user/saveProfile';
  _suggestions = 'suggestions/:id';

  _collections = 'collections/:id';
  _removeCollection = 'user/removeCollections';
  _msgs = 'msgs';

  login(email:string, password:string, rememberMe:string) {
    let that = this;
    return this._reqService.post(this._login, {email: email, password: password, rememberMe: rememberMe}).map((json:any) => {
      let errors = undefined;
      if (json.code == 1) {
        let days:number = rememberMe? 30: 1;

        that.saveProfileLocal(json.data, days);

        that.routeService.navTo('/pages/dashboard');
      } else {
        errors = json.msg;
      }
      return errors;
    });
  }
  logout() {
    this._reqService.post(this._logout, {}).subscribe((json:any) => {
      if (json.code == 1) {
        Cookie.delete(CONSTANT.PROFILE_KEY);

        this.routeService.navTo('/login');
      }
    });
  }
  register(name:string, phone:string, email:string, password:string) {
    let that = this;
    return this._reqService.post(this._register, {name:name, phone: phone, email: email, password: password}).map((json:any) => {
      let errors = undefined;
      if (json.code == 1) {
        that.saveProfileLocal(json.data, 1);

        that.routeService.navTo('/pages/dashboard');
      } else {
        errors = json.msg;
      }
      return errors;
    });
  }

  forgotPassword(email:number) {
    return this._reqService.post(this._forgotPassword, {email: email});
  }

  resetPassword(model:number) {
    let that = this;
    return this._reqService.post(this._resetPassword, model).map((json:any) => {
      let errors = undefined;
      if (json.code == 1) {
        that.saveProfileLocal(json.data, 1);

        that.routeService.navTo('/pages/dashboard');
      } else {
        errors = json.data;
      }
      return errors;
    });
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

  saveProfileLocal(profile:any, expireDays:number) {
    let that = this;
    CONSTANT.PROFILE = profile;

    if (!expireDays) {
      expireDays = parseInt(Cookie.get(CONSTANT.PROFILE_EXPIRE));
    } else {
      Cookie.set(CONSTANT.PROFILE_EXPIRE, expireDays + '', 365);
    }

    Cookie.set(CONSTANT.PROFILE_KEY, JSON.stringify(profile), expireDays);
    that._state.notifyDataChanged('profile.refresh', profile);
  }
  loadProfileLocal() {
    let that = this;
    let profile = Cookie.get(CONSTANT.PROFILE_KEY);

    if (profile) {
      CONSTANT.PROFILE = JSON.parse(profile);
      that._state.notifyDataChanged('profile.refresh', profile);
    }
  }
}
