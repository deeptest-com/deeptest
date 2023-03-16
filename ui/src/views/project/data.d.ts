export interface Project {
  id: number;
  name: string;
  desc: string;
  logo:string
  shortName:string
  adminId:number
  includeExample:bool
}

export interface QueryResult {
  list: Project[];
  pagination: PaginationConfig;
}

export interface QueryParams {
  keywords:  string,
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
