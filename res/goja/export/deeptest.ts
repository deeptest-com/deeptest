"use strict";

// datapool
const datapool = {} as any;

declare function getDatapoolVariable(variable_name: string): any;
datapool.get = getDatapoolVariable;
datapool.get.prototype = {};

// variables
const variables = {} as any;

declare function getVariable(variable_name: string): any;
variables.get = getVariable;
variables.get.prototype = {};

declare function setVariable(variable_name: string, variable_value: any): void;
variables.set = setVariable;
variables.set.prototype = {};

declare function clearVariable(variable_name: string): void;
variables.clear = clearVariable;
variables.clear.prototype = {};

export default {
    datapool,
    variables,
}