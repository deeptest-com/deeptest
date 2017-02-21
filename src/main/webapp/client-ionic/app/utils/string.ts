import {Injectable} from '@angular/core';

import {CONSTANT} from './constant';

@Injectable()
export class StringUtil {
    constructor() { }
    
    static IsString (str){
       return (typeof str == 'string') && str.constructor == String; 
    }
    
    
    static Trim (o){
        if (StringUtil.IsEmpty(o)) {
            return '';
        }
        
        o = o.replace(/(^\s*)|(\s*$)/g, '');
        return o;
    }
    
    static IsEmpty (o){
        if (o === null || o === "null" || o === undefined || o === "undefined" || o === "") {
            return true;
        } else {
            return false;
        }
    }
    
    static UpcaseFirst (str) {
        var first = str.substring(0,1).toUpperCase();
        var others = str.substring(1,str.length);
        var ret = first + others;
        return ret;
    }
}