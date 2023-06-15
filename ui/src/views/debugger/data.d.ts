export interface TestInterface {
  id: number;
  [props: any]: any;
}

export interface QueryResult {
  list: TestInterface[];
  pagination: PaginationConfig;
}

export interface QueryParams {
  keywords:  string,
  status: any,
  page: number,
  pageSize: number,
}

export interface PaginationConfig {
  total: number;
  current: number;
  pageSize: number;
  showSizeChanger: boolean;
  showQuickJumper: boolean;
}
