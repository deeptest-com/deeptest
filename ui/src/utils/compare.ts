import {ComparisonOperator} from "./enum";

export const getCompareOpts = () => {
  const arr : any[] = []

  arr.push({label: ComparisonOperator.contain, value: ComparisonOperator[ComparisonOperator.contain]})
  arr.push({label: ComparisonOperator.notContain, value: ComparisonOperator[ComparisonOperator.notContain]})

  arr.push({label: ComparisonOperator.equal, value: ComparisonOperator[ComparisonOperator.equal]})
  arr.push({label: ComparisonOperator.notEqual, value: ComparisonOperator[ComparisonOperator.notEqual]})
  arr.push({label: ComparisonOperator.greaterThan, value: ComparisonOperator[ComparisonOperator.greaterThan]})
  arr.push({label: ComparisonOperator.greaterThanOrEqual, value: ComparisonOperator[ComparisonOperator.greaterThanOrEqual]})
  arr.push({label: ComparisonOperator.lessThan, value: ComparisonOperator[ComparisonOperator.lessThan]})
  arr.push({label: ComparisonOperator.lessThanOrEqual, value: ComparisonOperator[ComparisonOperator.lessThanOrEqual]})

  return arr
}

export const getCompareOptsForRespCode = () => {
  const arr : any[] = []

  arr.push({label: ComparisonOperator.equal, value: ComparisonOperator[ComparisonOperator.equal]})
  arr.push({label: ComparisonOperator.notEqual, value: ComparisonOperator[ComparisonOperator.notEqual]})

  return arr
}

export const getCompareOptsForString = () => {
  const arr : any[] = []

  arr.push({label: ComparisonOperator.contain, value: ComparisonOperator[ComparisonOperator.contain]})
  arr.push({label: ComparisonOperator.notContain, value: ComparisonOperator[ComparisonOperator.notContain]})
  arr.push({label: ComparisonOperator.equal, value: ComparisonOperator[ComparisonOperator.equal]})
  arr.push({label: ComparisonOperator.notEqual, value: ComparisonOperator[ComparisonOperator.notEqual]})

  return arr
}
