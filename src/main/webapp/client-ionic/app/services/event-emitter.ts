import {Injectable} from '@angular/core';
import {Observable} from 'rxjs/Observable';
import {Subject} from 'rxjs/Subject';

@Injectable()
export class ChangeCategoryEventEmitter extends Subject<any> {
    constructor() {
        super();
    }
    emit(value) {
        super.next(value);
    }
}

@Injectable()
export class GotoTabEventEmitter extends Subject<number> {
    constructor() {
        super();
    }
    emit(value) {
        super.next(value);
    }
}

@Injectable()
export class ShoppingCartChangeEventEmitter extends Subject<number> {
    constructor() {
        super();
    }
    emit(value) {
        super.next(value);
    }
}
