

export interface QueryResult {
  list: Home[];
  pagination: PaginationConfig;
}

export interface QueryParams {
  keywords:  string,
  currProjectId:number,
  userId:number,
  enabled?: string,
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

