export interface Report {
  id: number;
  name: string;
  desc: string;
  logs?: any;
}

export interface QueryResult {
  list: Report[];
  pagination: PaginationConfig;
}

export interface QueryParams {
  keywords:  string,
  planId?: string,
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
