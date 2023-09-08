"use strict";

// datapool
const datapool = {};


datapool.get = function getDatapoolVariable(datapool_name, variable_name, seq) {};
datapool.get.prototype = {};

// variables
const variables = {};

variables.get = function getVariable(variable_name){};
variables.get.prototype = {};

variables.set = function setVariable(variable_name, variable_value){};
variables.set.prototype = {};

variables.clear = function clearVariable(variable_name){};
variables.clear.prototype = {};

module.exports = {
    datapool,
    variables,
}