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

module.exports = { // under dt.
    datapool,
    variables,

    test: test,
    expect: expect,
    sendRequest: sendRequest,
}
