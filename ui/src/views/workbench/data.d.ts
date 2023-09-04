

export interface QueryResult {
  list: Home[];
  pagination: PaginationConfig;
}

export interface QueryParams {

  projectId:number,
  cycle:number,//0代表按7天，1代表按30天
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

