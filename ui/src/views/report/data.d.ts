export interface Report {
  id: number;
  name: string;
  desc: string;
}

export interface QueryResult {
  list: Report[];
  pagination: PaginationConfig;
}

export interface QueryParams {
  keywords:  string,
  scenarioId: string,
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
