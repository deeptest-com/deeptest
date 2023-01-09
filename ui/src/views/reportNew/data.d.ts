export interface Report {
  id: number;
  name: string;
  desc: string;
  passRequestNum: number;
  failRequestNum: number;
  totalRequestNum: number;
  logs: ReportLog[];
}



export interface ReportLog{
  id: number;
  name: string;
  resultStatus: string;
  logs: ReportLogDetail[];

}

export interface ReportLogDetail{
  name: string;
  resultStatus: string;
  exectime: number;
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

// 分页设置
export interface PaginationConfig {
  total: number;
  current: number;
  pageSize: number;
  showSizeChanger: boolean;
  showQuickJumper: boolean;
}
