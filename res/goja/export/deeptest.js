"use strict";

// datapool
const datapool = {};

datapool.get = getDatapoolVariable;
datapool.get.prototype = {};

// variables
const variables = {};

variables.get = getVariable;
variables.get.prototype = {};

variables.set = setVariable;
variables.set.prototype = {};

variables.clear = clearVariable;
variables.clear.prototype = {};

// environment
// const environment = {};
//
// environment.get = getEnvironmentVariable;
// environment.get.prototype = {};
//
// environment.set = setEnvironmentVariable;
// environment.set.prototype = {};
//
// environment.clear = clearEnvironmentVariable;
// environment.clear.prototype = {};

const request = {}
const response = {}

module.exports = {
    datapool,
    variables,
    // environment,

    request,
    response,
}