import {Injectable} from '@angular/core';
import {VARIABLE} from '../utils/variable';
import {PostService} from './post';

@Injectable()
export class SearchService {
    constructor(private _postService: PostService) { }

    _search = '/search/query';
    _history = '/search/getHistory';
    _keywords = '/search/getMatchedKeywords';

    search(keywords) {
        return this._postService.post(this._search, {keywords: keywords, categoryId: VARIABLE.CURRENT_CATEGORY['id']});
    }

    getHistory() {
        return this._postService.post(this._history, {});
    }

    getMatchedKeywords(keywords) {
        return this._postService.post(this._keywords, {keywords: keywords, categoryId: VARIABLE.CURRENT_CATEGORY['id']});
    }
}
