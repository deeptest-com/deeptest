"use strict";

const variable = {};

const get = variable.get = getVariable;
get.prototype = {};

const set = variable.set = setVariable;
set.prototype = {};

module.exports = variable
