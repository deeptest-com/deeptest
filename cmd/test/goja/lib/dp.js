"use strict";

// environment
const environment = {};

environment.get = getEnvironmentVariable;
environment.get.prototype = {};

environment.set = setEnvironmentVariable;
environment.set.prototype = {};

environment.clear = clearEnvironmentVariable;
environment.clear.prototype = {};

// variables
const variables = {};

variables.get = getVariable;
variables.get.prototype = {};

variables.set = setVariable;
variables.set.prototype = {};

variables.clear = clearVariable;
variables.clear.prototype = {};

module.exports = {
    variables: variables,
    environment: environment,
}