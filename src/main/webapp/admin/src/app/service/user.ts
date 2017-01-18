import {Injectable} from "@angular/core";
import {RequestService} from "./request";

@Injectable()
export class UserService {
  constructor(private _reqService:RequestService) {
  }

  _login = 'user/login';
  _register = 'user/register';
  _forgotPassword = 'user/forgotPassword';
  _resetPassword = 'user/resetPassword';

  _getProfile = 'user/profile/get';
  _saveProfile = 'user/profile/save';
  _suggestions = 'suggestions/:id';

  _collections = 'collections/:id';
  _removeCollection = 'user/removeCollections';
  _msgs = 'msgs';


  login(email, password, rememberMe) {
    return this._reqService.post(this._login, {email: email, password: password, rememberMe: rememberMe});
  }

  register(name, email, password) {
    return this._reqService.post(this._register, {name:name, email: email, password: password});
  }

  forgotPassword(phone) {
    return this._reqService.post(this._forgotPassword, {phone: phone});
  }

  resetPassword(phone) {
    return this._reqService.post(this._resetPassword, {phone: phone});
  }

  getProfile() {
    return this._reqService.post(this._getProfile, {});
  }

  saveProfile(profile) {
    return this._reqService.post(this._saveProfile, profile);
  }

  saveSuggestion(content) {
    return this._reqService.post(this._suggestions.replace(':id', ''), {suggestion: {content: content}});
  }
}

