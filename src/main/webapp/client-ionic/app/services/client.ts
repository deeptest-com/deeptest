import {Injectable} from '@angular/core';
import {PostService} from './post';

@Injectable()
export class ClientService {
    constructor(private _postService: PostService) { }

    _getProfile = '/client/profile/get';
    _saveProfile = '/client/profile/save';
    _forgotPassword = '/client/forgotPassword';
    _signon = '/client/signon';
    _signup = '/client/signup';

    _suggestions = '/suggestions/:id';
    _resetPassword = '/client/resetPassword';

    _collections = '/collections/:id';
    _removeCollection = '/client/removeCollections';
    _msgs = '/msgs';

    getProfile() {
        return this._postService.post(this._getProfile, {});
    }

    saveProfile(profile) {
        return this._postService.post(this._saveProfile, profile);
    }

    forgotPassword(phone) {
        return this._postService.post(this._forgotPassword, {phone:phone});
    }

    signon(phone, password, rememberMe) {
        return this._postService.post(this._signon, {phone:phone, password:password, rememberMe:rememberMe});
    }

    signup(phone, password, repassword) {
        return this._postService.post(this._signup, {phone:phone, password:password, repassword:repassword});
    }

    saveSuggestion(content) {
        return this._postService.post(this._suggestions.replace(':id', ''), {suggestion: {content: content}});
    }

    resetPassword(phone) {
        return this._postService.post(this._resetPassword, {phone: phone});
    }
}
