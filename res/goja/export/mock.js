"use strict";

const getParam = function getVariable(name){
    var p = dt.request.queryParams.find(i => i.name === name)
    if (!p) return ''
    return p.value
};
getParam.prototype = {};

const getHeader = function getVariable(name){
    var p = dt.request.headers.find(i => i.name === name.toUpperCase())
    if (!p) return ''
    return p.value
};
getHeader.prototype = {};

const getCookie = function getVariable(name){
    var p = dt.request.cookies.find(i => i.name === name)
    if (!p) return null

    return p
};
getCookie.prototype = {};

module.exports = {
    getParam, getHeader, getCookie
}